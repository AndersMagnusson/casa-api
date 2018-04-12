package vapix

import (
	"casa/src/server/pkg/httpnet"
	"fmt"
)

var protocol = "http"

const (
	serialNumberPath = "axis-cgi/operator/param.cgi?action=list&group=Properties.System.SerialNumber"
)

func CheckCredentials(address string, username string, password string) (int, error) {
	url := fmt.Sprintf("%s://%s/%s", protocol, address, serialNumberPath)
	resp, err := httpnet.DoBasic(url, "GET", nil, nil, httpnet.Credentials{Username: username, Password: password})
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("Failed with body: %s", resp.Body)
	}
	return resp.StatusCode, nil
}
