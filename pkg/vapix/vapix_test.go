package vapix

import (
	"encoding/xml"
	"testing"
)

func TestApplicationXml(t *testing.T) {

	var response ApplicationsXml
	err := xml.Unmarshal([]byte(`
		<reply result="ok">
 <application Name="RemoteAccess" NiceName="AXIS Remote Access solution" Vendor="Axis Communications" Version="1.14" ApplicationID="1234" License="Custom" Status="Running" ConfigurationPage="local/RemoteAccess/#" VendorHomePage="http://www.axis.com" />
 <application Name="VMD3" NiceName="AXIS Video Motion Detection" Vendor="Axis Communications" Version="3.1-1" ApplicationID="46396" License="None" Status="Running" ConfigurationPage="local/VMD3/setup.html" VendorHomePage="http://www.axis.com" />
 <application Name="deviceDiagnostics" NiceName="AXIS Device Diagnostics" Vendor="Axis Communications" Version="1.26690" ApplicationID="328106" License="None" Status="Running" ConfigurationPage="local/deviceDiagnostics/info.html" />
 <application Name="vmd" NiceName="AXIS Video Motion Detection" Vendor="Axis Communications" Version="4.2-4" ApplicationID="143440" License="None" Status="Stopped" ConfigurationPage="local/vmd/config.html" VendorHomePage="http://www.axis.com" />
</reply>`), &response)
	if err != nil {
		t.Fatal("Failed to unmarshal", err)
	}
	if response.Result != "ok" {
		t.Fatal("Failed to get result")
	}
	if len(response.Applications) != 4 {
		t.Fatalf("Should get %d but got %d", 5, len(response.Applications))
	}
	if response.Applications[0].Name != "RemoteAccess" {
		t.Fatalf("Expected name to be %s but was %s", "AxisConnectDeploymentAgentAcap", response.Applications[0].Name)
	}
	if response.Applications[0].Status != "Running" {
		t.Fatalf("Expected name to be %s but was %s", "Running", response.Applications[0].Status)
	}
	if response.Applications[0].Version != "1.14" {
		t.Fatalf("Expected name to be %s but was %s", "1.14", response.Applications[0].Version)
	}

}
