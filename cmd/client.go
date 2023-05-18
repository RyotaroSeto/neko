package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type HttpClient struct {
	url           string
	method        string
	requestHeader map[string]string
}

func NewHttpClient(url string, method string) *HttpClient {
	var hc HttpClient
	hc.requestHeader = map[string]string{
		"Connection": "keep-alive",
	}
	return &HttpClient{
		url:           url,
		method:        method,
		requestHeader: hc.requestHeader,
	}
}

func (c *HttpClient) setRequestHeader(req *http.Request) {
	for k, v := range c.requestHeader {
		req.Header.Set(k, v)
	}
}

func (c *HttpClient) sendRequest() (*http.Response, error) {
	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		return nil, err
	}

	c.setRequestHeader(req)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type APIResponse struct {
	Err  int        `json:"err"`
	Name [][]string `json:"name"`
}

func (c *HttpClient) Execute() (*APIResponse, error) {
	res, err := c.sendRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`callback\((.*)\)`)
	matches := re.FindSubmatch(body)
	if len(matches) != 2 {
		fmt.Println("JSONPレスポンスのパースエラー")
		return nil, err
	}

	var apiResponse APIResponse
	err = json.Unmarshal(matches[1], &apiResponse)
	if err != nil {
		fmt.Println("JSONパースエラー:", err)
		return nil, err
	}

	if apiResponse.Err != 0 {
		fmt.Println("APIエラー:", apiResponse.Err)
		return nil, err
	}

	return &apiResponse, nil
}
