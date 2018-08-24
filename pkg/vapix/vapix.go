package vapix

import (
	"casa-api/pkg/httpnet"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

var protocol = "http"

const (
	EmbeddedDevelopmentVersionPath = "Properties.EmbeddedDevelopment.Version"
	serialNumberPath               = "axis-cgi/operator/param.cgi?action=list&group=Properties.System.SerialNumber"
	ListApplications               = "/axis-cgi/applications/list.cgi"
)

func CheckCredentials(address string, username string, password string) (int, error) {
	url := fmt.Sprintf("%s://%s/%s", protocol, address, serialNumberPath)
	resp, err := httpnet.DoDigest(url, "GET", nil, nil, httpnet.Credentials{Username: username, Password: password})
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("Failed with body: %s", resp.Body)
	}
	return resp.StatusCode, nil
}

func GetProperty(address, username, password, property string) (string, error) {
	url := fmt.Sprintf("%s://%s/%s", protocol, address, serialNumberPath)
	resp, err := httpnet.DoDigest(url, "GET", nil, nil, httpnet.Credentials{Username: username, Password: password})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Failed with body: %s", resp.Body)
	}
	return resp.Body, nil
}

func GetApplications(address, username, password string) (string, error) {
	url := fmt.Sprintf("%s://%s/%s", protocol, address, ListApplications)
	resp, err := httpnet.DoDigest(url, "GET", nil, nil, httpnet.Credentials{Username: username, Password: password})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Failed with body: %s", resp.Body)
	}
	return resp.Body, nil
}

// GetHighestVmdVersion returns the highest major vmd version. Defaults to version 2.
func GetHighestVmdVersion(address, username, password string) (int, error) {
	data, err := GetApplications(address, username, password)
	if err != nil {
		return 0, err
	}
	var applications ApplicationsXml
	err = xml.Unmarshal([]byte(data), &applications)
	if err != nil {
		return 0, err
	}
	version := 2 // default to 2
	for _, a := range applications.Applications {
		if a.NiceName == "AXIS Video Motion Detection" && a.Status == "Running" {
			v := strings.Split(a.Version, ".")
			if len(v) > 0 {
				vv, err := strconv.Atoi(v[0])
				if err == nil && vv > version {
					version = vv
				}
			}
		}
	}
	return version, nil
}

// <reply result="ok">
//  <application Name="AxisConnectDeploymentAgentAcap" NiceName="AXIS Connect Deployment Agent" Vendor="Axis Communications" Version="0.0-1835" License="None" Status="Running" ConfigurationPage="local/AxisConnectDeploymentAgentAcap/index.html" />
//  <application Name="RemoteAccess" NiceName="AXIS Remote Access solution" Vendor="Axis Communications" Version="1.14" ApplicationID="1234" License="Custom" Status="Running" ConfigurationPage="local/RemoteAccess/#" VendorHomePage="http://www.axis.com" />
//  <application Name="VMD3" NiceName="AXIS Video Motion Detection" Vendor="Axis Communications" Version="3.1-1" ApplicationID="46396" License="None" Status="Running" ConfigurationPage="local/VMD3/setup.html" VendorHomePage="http://www.axis.com" />
//  <application Name="deviceDiagnostics" NiceName="AXIS Device Diagnostics" Vendor="Axis Communications" Version="1.26690" ApplicationID="328106" License="None" Status="Running" ConfigurationPage="local/deviceDiagnostics/info.html" />
//  <application Name="vmd" NiceName="AXIS Video Motion Detection" Vendor="Axis Communications" Version="4.2-4" ApplicationID="143440" License="None" Status="Stopped" ConfigurationPage="local/vmd/config.html" VendorHomePage="http://www.axis.com" />
// </reply>
type ApplicationsXml struct {
	XMLName      xml.Name      `xml:"reply"`
	Result       string        `xml:"result,attr"`
	Applications []Application `xml:"application"`
}

type Application struct {
	XMLName           xml.Name `xml:"application"`
	Name              string   `xml:"Name,attr"`
	NiceName          string   `xml:"NiceName,attr"`
	Version           string   `xml:"Version,attr"`
	Vendor            string   `xml:"Vendor,attr"`
	License           string   `xml:"License,attr"`
	Status            string   `xml:"Status,attr"`
	ConfigurationPage string   `xml:"ConfigurationPage,attr"`
}
