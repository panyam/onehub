package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	gut "github.com/panyam/goutils/utils"
)

var ErrEntityNotFound = errors.New("entity not found")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrAuthorized = errors.New("unauthorized")
var ErrInvalidRequest = errors.New("request is invalid")

type TSClient struct {
	Host       string
	ApiKey     string
	LogRequest bool
	LogBody    bool
}

func NewClient(host string, apikey string) *TSClient {
	if strings.TrimSpace(host) == "" {
		host = os.Getenv("TYPESENSE_HOST")
		if strings.TrimSpace(host) == "" {
			host = "http://localhost:8108"
		}
	}
	if strings.TrimSpace(apikey) == "" {
		apikey = os.Getenv("TYPESENSE_API_KEY")
		if strings.TrimSpace(apikey) == "" {
			apikey = "test_api_key"
		}
	}
	return &TSClient{
		Host:       host,
		ApiKey:     apikey,
		LogRequest: true,
		LogBody:    true,
	}
}

func (t *TSClient) Request(method string, endpoint string, body gut.StringMap) (gut.StringMap, error) {
	return t.RequestWithArgs(method, endpoint, "", body)
}

func (t *TSClient) RequestWithArgs(method string, endpoint string, args string, body gut.StringMap) (response gut.StringMap, err error) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	var req *http.Request
	var resp *http.Response
	url := fmt.Sprintf("%s/%s", t.Host, endpoint)
	if args != "" {
		url += "?" + args
	}
	var reqBody []byte
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader := bytes.NewBuffer(reqBody)
		req, err = http.NewRequest(method, url, bodyReader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return
	}

	req.Header.Set("X-TYPESENSE-API-KEY", t.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// create insecure transport

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	log.Printf("Request: '%s %s", method, url)
	if body != nil {
		marshalled, _ := json.MarshalIndent(body, "", "  ")
		log.Println("BODY: ", string(marshalled))
	}
	startTime := time.Now()
	resp, err = client.Do(req)
	if err != nil {
		log.Println("client: error making http request: ", err)
		return nil, err
	}
	endTime := time.Now()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Response: %d in %f seconds", resp.StatusCode, (endTime.Sub(startTime)).Seconds())
	if resp.StatusCode != 200 {
		log.Println("Response Message: ", string(respbody))
	}

	if resp.StatusCode == 400 {
		return nil, ErrInvalidRequest
	} else if resp.StatusCode == 401 {
		return nil, ErrInvalidCredentials
	} else if resp.StatusCode == 403 {
		return nil, ErrAuthorized
	} else if resp.StatusCode == 404 {
		return nil, ErrEntityNotFound
	}

	err = json.Unmarshal(respbody, &response)
	return response, err
}

/**
 * Gets a collection schema.
 */
func (t *TSClient) GetCollection(doctype string) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	return t.Request("GET", endpoint, nil)
}

func (t *TSClient) UpdateCollection(doctype string, fields []gut.StringMap) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	return t.Request("PATCH", endpoint, gut.StringMap{
		"fields": fields,
	})
}

/**
 * Create a collection with the given field info.
 */
func (t *TSClient) CreateCollection(schema gut.StringMap) (gut.StringMap, error) {
	endpoint := "/collections"
	return t.Request("POST", endpoint, schema)
}

/**
 * Delete a collection.
 */
func (t *TSClient) DeleteCollection(doctype string) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	return t.Request("DELETE", endpoint, nil)
}

/**
 * Gets a document by its ID in a given collection.
 */
func (t *TSClient) GetDocument(doctype string, docid string) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	return t.Request("GET", endpoint, nil)
}

/**
 * Gets a document by its ID in a given collection.
 */
func (t *TSClient) DeleteDocument(doctype string, docid string) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	return t.Request("DELETE", endpoint, nil)
}

/**
 * Upserts a document given its ID into a collection.
 */
func (t *TSClient) Upsert(doctype string, docid string, doc gut.StringMap) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	return t.Request("PATCH", endpoint, doc)
}

func fieldDifferent(oldField gut.StringMap, newField gut.StringMap) bool {
	if oldField == nil {
		return true
	}
	newFieldType := newField["type"].(string)
	newFieldOptional := false
	if _, ok := newField["optional"]; ok {
		newFieldOptional = newField["optional"].(bool)
	}
	oldFieldType := oldField["type"].(string)
	oldFieldOptional := false
	if _, ok := oldField["optional"]; ok {
		oldFieldOptional = oldField["optional"].(bool)
	}
	if newFieldType != oldFieldType || oldFieldOptional != newFieldOptional {
		return true
	}
	// TODO - check other field attribs too
	return false
}

func (t *TSClient) EnsureSchema(doctype string, newFields []gut.StringMap) {
	// fields, fieldMap := PGTableInfoToSchema(tableInfo)
	existing, err := t.GetCollection(doctype)
	if err != nil {
		log.Println("Schema Fetch Error: ", err)
	}
	if existing == nil {
		schema := gut.StringMap{
			"name":                 doctype,
			"enable_nested_fields": true,
			"fields":               newFields,
		}
		res, err := t.CreateCollection(schema)
		log.Println("Schema Creation: ", doctype, res, err)
		if err != nil {
			panic(err)
		}
	} else {
		oldFields := existing["fields"].([]interface{})
		oldFieldMap := make(map[string]gut.StringMap)
		for _, f := range oldFields {
			field := f.(gut.StringMap)
			name := field["name"].(string)
			oldFieldMap[name] = field
		}
		newFieldMap := make(map[string]gut.StringMap)
		for _, field := range newFields {
			name := field["name"].(string)
			newFieldMap[name] = field
		}
		// Go through *new* fields and see which either dont exist
		// or have changed and add those
		var patchinfo []gut.StringMap
		for _, newField := range newFields {
			newFieldName := newField["name"].(string)
			// Does it not exist in prev?
			if newFieldName == "id" {
				// Field `id` cannot be altered.
				continue
			}
			oldField, ok := oldFieldMap[newFieldName]
			if !ok || oldField == nil {
				// It does not exist in the old set so just add it
				patchinfo = append(patchinfo, newField)
			} else {
				if fieldDifferent(newField, oldField) {
					// drop and reload it
					patchinfo = append(patchinfo, gut.StringMap{
						"drop": true,
						"name": newFieldName,
					})

					// now added
					patchinfo = append(patchinfo, newField)
				}
			}
		}

		// Now go through fields that have been "droped"
		for _, f := range oldFields {
			oldField := f.(gut.StringMap)
			oldFieldName := oldField["name"].(string)
			// Does it not exist in prev?
			if oldFieldName == "id" {
				// Field `id` cannot be altered.
				continue
			}
			newField, ok := newFieldMap[oldFieldName]
			if !ok || newField == nil {
				// Does not exist in new so drop it
				// It does not exist in the old set so just add it
				patchinfo = append(patchinfo, gut.StringMap{
					"drop": true,
					"name": oldFieldName,
				})
			}
		}

		if patchinfo != nil {
			res, err := t.UpdateCollection(doctype, patchinfo)
			log.Println("Schema Update: ", doctype, res, err)
			if err != nil {
				panic(err)
			}
		}
	}
}
