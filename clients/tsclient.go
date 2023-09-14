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

type TSClient struct {
	Host   string
	ApiKey string
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
		Host:   host,
		ApiKey: apikey,
	}
}

func (t *TSClient) Request(method string, endpoint string, body gut.StringMap) (gut.StringMap, error) {
	return t.RequestWithArgs(method, endpoint, "", body)
}

func (t *TSClient) RequestWithArgs(method string, endpoint string, args string, body gut.StringMap) (gut.StringMap, error) {
	if strings.HasPrefix(endpoint, "/") {
		endpoint = endpoint[1:]
	}
	url := fmt.Sprintf("%s/%s", t.Host, endpoint)
	if args != "" {
		url += "?" + args
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewBuffer(reqBody)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}

	log.Printf("Sending request: 'POST %s", url)
	log.Printf("BODY: %s", reqBody)

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
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response gut.StringMap
	err = json.Unmarshal(respbody, &response)
	return response, err
}

func (t *TSClient) Upsert(doctype string, docid string, doc gut.StringMap) (out gut.StringMap, err error) {
	endpoint := fmt.Sprintf("collections/%s/documents/%s", doctype, docid)
	return t.Request("PATCH", endpoint, doc)
}
