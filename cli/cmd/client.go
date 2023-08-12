package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type StringMap = map[string]interface{}

type OHClient struct {
	Host     string
	Username string
	Password string

	PrettyPrintResponse bool
	transport           *http.Transport
}

func NewOHClient(host string) *OHClient {
	// remove the trailing slash
	if strings.HasSuffix(host, "/") {
		host = host[:len(host)-1]
	}

	return &OHClient{
		Host:                host,
		PrettyPrintResponse: true,
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// Make a JSON POST request to the OHClient
func (c *OHClient) Call(method string, path string, args map[string][]string, headers StringMap, body StringMap) (response StringMap, err error) {
	url := fmt.Sprintf("%s%s", c.Host, path)
	var req *http.Request

	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		reqBody, err := json.Marshal(body)
		if err == nil {
			bodyReader := bytes.NewBuffer(reqBody)
			req, err = http.NewRequest(method, url, bodyReader)
		}
	}
	if err != nil {
		return nil, err
	}

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
	if err != nil {
		fmt.Println(err)
	} else if c.PrettyPrintResponse {
		indented, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println(string(indented))
	}
	return
}
