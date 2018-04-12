package httpnet

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}

type Response struct {
	Body       string
	StatusCode int
}

type Credentials struct {
	Password string
	Username string
}

func Do(url string, method string, postData interface{}, respData interface{}) (Response, error) {
	return do(url, method, postData, respData, Credentials{})
}

func DoBasic(url string, method string, postData interface{}, respData interface{}, credentials Credentials) (Response, error) {
	return do(url, method, postData, respData, credentials)
}

func do(url string, method string, postData interface{}, respData interface{}, credentials Credentials) (Response, error) {
	d := Response{StatusCode: -1}
	var buffer io.Reader
	if postData != nil {
		jsonBytes, err := json.Marshal(postData)
		if err != nil {
			return d, err
		}
		buffer = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(method, url, buffer)
	if buffer != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if len(credentials.Password) > 0 || len(credentials.Username) > 0 {
		req.SetBasicAuth(credentials.Username, credentials.Password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}

	if respData != nil {
		if len(body) > 0 {
			err = json.Unmarshal(body, respData)
			if err != nil {
				return d, err
			}
		}
	}
	d.StatusCode = resp.StatusCode
	d.Body = string(body)
	return d, nil
}
