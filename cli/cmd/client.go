package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type OHClient struct {
	Host     string
	Username string
	Password string

	transport *http.Transport
}

func NewOHClient(host string) *OHClient {
	// remove the trailing slash
	if strings.HasSuffix(host, "/") {
		host = host[:len(host)-1]
	}

	return &OHClient{
		Host: host,
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// Make a JSON POST request to the OHClient
func (c *OHClient) Call(method string, path string, args map[string][]string, body map[string]interface{}, headers map[string]interface{}) (response map[string]interface{}, err error) {
	// make a post reuqest
	url := fmt.Sprintf("%s%s", c.Host, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Set the body
	if body != nil {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader := bytes.NewBuffer(reqBody)
		req, err = http.NewRequest(method, url, bodyReader)
		if err != nil {
			return nil, err
		}
		log.Printf("BODY: %s", reqBody)
	} else {
	}

	log.Printf("Sending request: '%s %s", method, url)

	// Set any custom headers needed
	if headers == nil {
		for hdrname, hdrvalue := range headers {
			req.Header.Set(hdrname, fmt.Sprintf("%v", hdrvalue))
		}
	}
	if c.Username != "" && c.Password != "" {
		up := fmt.Sprintf("%s:%s", c.Username, c.Password)
		b64 := base64.StdEncoding.EncodeToString([]byte(up))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64))
	}
	req.Header.Set("Content-Type", "application/json")

	// create a http client
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: c.transport,
	}

	// and make the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respbody, &response)
	return
}
