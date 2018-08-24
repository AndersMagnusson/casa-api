package httpnet

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bobziuchkovski/digest"
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

func DoDigest(url string, method string, postData interface{}, respData interface{}, credentials Credentials) (Response, error) {
	d := Response{StatusCode: -1}

	req, err := createRequest(url, method, postData)
	if err != nil {
		return d, err
	}
	// req.Header.Add("User-Agent", "Axis Companion/3.45.018")

	t := digest.NewTransport(credentials.Username, credentials.Password)
	resp, err := t.RoundTrip(req)

	return parseResponse(resp, respData)
}

func do(url string, method string, postData interface{}, respData interface{}, credentials Credentials) (Response, error) {
	d := Response{StatusCode: -1}

	req, err := createRequest(url, method, postData)
	if err != nil {
		return d, err
	}

	if len(credentials.Password) > 0 || len(credentials.Username) > 0 {
		req.SetBasicAuth(credentials.Username, credentials.Password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return d, err
	}

	return parseResponse(resp, respData)
}

func createRequest(url string, method string, postData interface{}) (*http.Request, error) {
	var buffer io.Reader
	if postData != nil {
		jsonBytes, err := json.Marshal(postData)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	if buffer != nil {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Axis-Orig-Sw", strconv.FormatBool(true))
	}
	return req, nil
}

func parseResponse(resp *http.Response, respData interface{}) (Response, error) {
	d := Response{StatusCode: -1}
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
