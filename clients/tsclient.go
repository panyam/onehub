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

func (t *TSClient) EnsureSchema(doctype string, fields []gut.StringMap) {
	fieldMap := make(map[string]gut.StringMap)
	for _, field := range fields {
		name := field["name"].(string)
		fieldMap[name] = field
	}
	// fields, fieldMap := PGTableInfoToSchema(tableInfo)
	schema := gut.StringMap{
		"name":                 doctype,
		"enable_nested_fields": true,
		"fields":               fields,
	}
	existing, err := t.GetCollection(doctype)
	if err != nil {
		log.Println("Schema Fetch Error: ", err)
	}
	if existing == nil {
		res, err := t.CreateCollection(schema)
		log.Println("Schema Creation: ", doctype, res, err)
		if err != nil {
			panic(err)
		}
	} else {
		// TODO - check there are *acutally* changes first
		// update it
		var newFields []gut.StringMap
		for _, ef := range existing["fields"].([]interface{}) {
			efield := ef.(gut.StringMap)
			fieldName := efield["name"].(string)
			if fieldName == "id" {
				// Field `id` cannot be altered.
				continue
			}
			fieldType := efield["type"].(string)
			fieldOptional := efield["optional"].(bool)
			newField := fieldMap[fieldName]
			newFieldName := newField["name"].(string)
			newFieldType := newField["type"].(string)
			newFieldOptional := false
			if _, ok := newField["optional"]; ok {
				newFieldOptional = newField["optional"].(bool)
			}
			if newFieldName != fieldName {
				// New field added
				newFields = append(newFields, newField)
			} else if newFieldType != fieldType || fieldOptional != newFieldOptional {
				// drop and reload it
				newFields = append(newFields, gut.StringMap{
					"drop": true,
					"name": newFieldName,
					"type": fieldType,
				})

				// now added
				newFields = append(newFields, gut.StringMap{
					"name":     newFieldName,
					"type":     newFieldType,
					"optional": true,
				})
			}
		}
		if newFields != nil {
			res, err := t.UpdateCollection(doctype, newFields)
			log.Println("Schema Update: ", doctype, res, err)
			if err != nil {
				panic(err)
			}
		}
	}
}
