package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	gut "github.com/panyam/goutils/utils"
)

const (
	ErrCodeInvalidRequest     = 400
	ErrCodeInvalidCredentials = 401
	ErrCodeAuthorized         = 403
	ErrCodeEntityNotFound     = 404
)

func TSErrorCode(err error) int {
	if err != nil {
		switch e := err.(type) {
		case *TSError:
			return e.Code
		default:
		}
	}
	return -1
}

type TSError struct {
	Code    int
	Message string
}

func (t *TSError) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", t.Code, t.Message)
}

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

func (t *TSClient) MakeUrl(endpoint string, args string) (url string) {
	endpoint = strings.TrimPrefix(endpoint, "/")
	url = fmt.Sprintf("%s/%s", t.Host, endpoint)
	if args != "" {
		url += "?" + args
	}
	return url
}

func (t *TSClient) MakeJsonRequest(method string, endpoint string, body gut.StringMap) (req *http.Request, err error) {
	var bodyBytes []byte
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	if body != nil {
		marshalled, _ := json.MarshalIndent(body, "", "  ")
		log.Println("BODY: ", string(marshalled))
	}
	return t.MakeBytesRequest(method, endpoint, bodyBytes)
}

func (t *TSClient) MakeBytesRequest(method string, endpoint string, body []byte) (req *http.Request, err error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}
	return t.MakeRequest(method, endpoint, bodyReader)
}

func (t *TSClient) MakeRequest(method string, endpoint string, bodyReader io.Reader) (req *http.Request, err error) {
	url := t.MakeUrl(endpoint, "")
	req, err = http.NewRequest(method, url, bodyReader)
	if err == nil {
		req.Header.Set("X-TYPESENSE-API-KEY", t.ApiKey)
		req.Header.Set("Content-Type", "application/json")
		log.Printf("Request: '%s %s", method, url)
	}
	return
}

func (t *TSClient) JsonCall(req *http.Request) (response gut.StringMap, err error) {
	out, err := t.Call(req)
	if err != nil {
		return nil, err
	}
	return out.(gut.StringMap), err
}

func (t *TSClient) Call(req *http.Request) (response interface{}, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	startTime := time.Now()
	resp, err := client.Do(req)
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

	if resp.StatusCode >= 400 {
		return nil, &TSError{resp.StatusCode, string(respbody)}
	}

	content_type := resp.Header.Get("Content-Type")
	if strings.HasPrefix(content_type, "application/json") {
		err = json.Unmarshal(respbody, &response)
	} else {
		// send as is
		response = respbody
	}
	return response, err
}

/**
 * Gets a collection schema.
 */
func (t *TSClient) GetCollection(doctype string) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	req, err := t.MakeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

func (t *TSClient) UpdateCollection(doctype string, fields []gut.StringMap) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	req, err := t.MakeJsonRequest("PATCH", endpoint, gut.StringMap{
		"fields": fields,
	})
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

/**
 * Create a collection with the given field info.
 */
func (t *TSClient) CreateCollection(schema gut.StringMap) (gut.StringMap, error) {
	endpoint := "/collections"
	req, err := t.MakeJsonRequest("POST", endpoint, schema)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

/**
 * Delete a collection.
 */
func (t *TSClient) DeleteCollection(doctype string) (gut.StringMap, error) {
	endpoint := fmt.Sprintf("collections/%s", doctype)
	req, err := t.MakeRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

/**
 * Gets a document by its ID in a given collection.
 */
func (t *TSClient) Get(doctype string, docid string) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	req, err := t.MakeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

/**
 * Gets a document by its ID in a given collection.
 */
func (t *TSClient) Delete(doctype string, docid string) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	req, err := t.MakeRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

/**
 * Upserts a document given its ID into a collection.
 */
func (t *TSClient) Upsert(doctype string, docid string, doc gut.StringMap) (out gut.StringMap, err error) {
	if len(strings.TrimSpace(docid)) > 0 {
		doc["id"] = strings.TrimSpace(docid)
	}
	endpoint := fmt.Sprintf("collections/%s/documents?action=upsert", doctype)
	req, err := t.MakeJsonRequest("POST", endpoint, doc)
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
}

func (t *TSClient) BatchUpsert(doctype string, docs []gut.StringMap) (interface{}, error) {
	var jsonlines []byte
	for i, d := range docs {
		marshalled, _ := json.Marshal(d)
		if i > 0 {
			jsonlines = append(jsonlines, '\n')
		}
		jsonlines = append(jsonlines, marshalled...)
	}
	endpoint := fmt.Sprintf("collections/%s/documents/import?action=upsert", doctype)
	req, err := t.MakeBytesRequest("POST", endpoint, jsonlines)
	req.Header.Set("Content-Type", "application/jsonlines")
	if err != nil {
		return nil, err
	}
	if out, err := t.Call(req); err != nil {
		return nil, err
	} else {
		payload := string(out.([]byte))
		jsonlines := strings.Split(payload, "\n")
		return gut.Map(jsonlines, func(input string) gut.StringMap {
			var out gut.StringMap
			json.Unmarshal([]byte(input), &out)
			return out
		}), nil
	}
}

func (t *TSClient) BatchDelete(doctype string, docids []string) (interface{}, error) {
	/*
		if false {
			result, err := t.DeleteDocument(doctype, docid)
			// result, err := tsclient.Collections(doctype).Documents(docid).Delete()
			if err != nil && err != ErrEntityNotFound {
				schema, err2 := tsclient.DeleteCollection(doctype)
				log.Println("Error Deleting: ", result, err)
				log.Println("Old Schema: ", schema, err2)
				panic(err)
			}
		}
	*/
	endpoint := fmt.Sprintf("collections/%s/documents?filter_by=id:[%s]", doctype, strings.Join(docids, ","))
	req, err := t.MakeRequest("DELETE", endpoint, nil)
	req.Header.Set("Content-Type", "application/jsonlines")
	if err != nil {
		return nil, err
	}
	return t.JsonCall(req)
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
