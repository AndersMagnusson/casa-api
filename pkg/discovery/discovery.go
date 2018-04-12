package discovery

import (
	"casa/src/server/pkg/httpnet"
	"casa/src/server/pkg/status"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/fromkeith/gossdp"
)

var discoveryRunning bool
var discoveryRunningLock sync.Mutex
var IsDiscoveryRunning sync.WaitGroup

type clientListener struct {
}

type AxisDevice struct {
	SerialNumber   string
	ModelName      string
	ModelNumber    string
	Address        string
	LastDiscovered time.Time
}

var discoveredDevicesLock sync.Mutex
var discoveredDevices = make(map[string]AxisDevice)

func GetDiscoveredDevices() map[string]AxisDevice {
	discoveredDevicesLock.Lock()
	defer discoveredDevicesLock.Unlock()
	result := make(map[string]AxisDevice)
	for key, value := range discoveredDevices {
		result[key] = value
	}
	return result
}

func GetDiscoveredDevice(serialNumber string) (AxisDevice, bool) {
	discoveredDevicesLock.Lock()
	device, ok := discoveredDevices[serialNumber]
	discoveredDevicesLock.Unlock()
	return device, ok
}

func (l *clientListener) Response(message gossdp.ResponseMessage) {
	go func(location string) {
		if len(location) > 0 {
			fmt.Printf("location: %s", location)
			resp, err := httpnet.DoBasic(location, "GET", nil, nil, httpnet.Credentials{})
			if err != nil {
				return
			}
			if len(resp.Body) > 0 {
				var ssdpInfo root
				err = xml.Unmarshal([]byte(resp.Body), &ssdpInfo)
				if err != nil {
					fmt.Printf("Unmarshal error: %s, %s", location, err.Error())
					return
				}
				if strings.ToLower(ssdpInfo.Manufacturer) == "axis" {
					host := ""
					url, err := url.Parse(ssdpInfo.PresentationURL)
					if err == nil {
						host = url.Hostname()
					}
					device := AxisDevice{
						LastDiscovered: time.Now(),
						ModelName:      ssdpInfo.ModelName,
						ModelNumber:    ssdpInfo.ModelNumber,
						SerialNumber:   ssdpInfo.SerialNumber,
						Address:        host,
					}
					discoveredDevicesLock.Lock()
					discoveredDevices[ssdpInfo.SerialNumber] = device
					discoveredDevicesLock.Unlock()
					status.UpdateDiscoveryStatus(device.SerialNumber, device.Address)
				}
			}
		}
	}(message.Location)
}

func SetCredential(serialNumber, username, password string) (int, error) {
	var axisDevice AxisDevice
	discoveredDevicesLock.Lock()
	axisDevice, ok := discoveredDevices[serialNumber]
	discoveredDevicesLock.Unlock()

	if ok {
		currentStatus := status.UpdateStatus(serialNumber, axisDevice.Address, username, password)
		return currentStatus.StatusCode, nil
	}
	return 404, nil
}

func Start() {
	discoveryRunningLock.Lock()
	if discoveryRunning {
		discoveryRunningLock.Unlock()
		return
	}

	IsDiscoveryRunning.Add(1)
	discoveryRunning = true
	discoveryRunningLock.Unlock()

	defer func() {
		discoveryRunningLock.Lock()
		discoveryRunning = false
		discoveryRunningLock.Unlock()
	}()

	listener := &clientListener{}
	c, err := gossdp.NewSsdpClient(listener)
	if err != nil {
		log.Println("Failed to start client: ", err)
		return
	}
	defer c.Stop()
	go c.Start()

	c.ListenFor("urn:schemas-upnp-org:device:Basic:1")
	// err = c.ListenFor("urn:fromkeith:test:web:1")
	time.Sleep(30 * time.Second)
	IsDiscoveryRunning.Done()
}

func DiscoveryRoutine() {
	go func() {
		Start()
		ticker := time.NewTicker(time.Minute * 60)
		for {
			select {
			case <-ticker.C:
				Start()
			}

		}
	}()
}

type root struct {
	Major            string `xml:"specVersion>major"`
	ManufacturerURL  string `xml:"device>manufacturerURL"`
	ModelDescription string `xml:"device>modelDescription"`
	ModelName        string `xml:"device>modelName"`
	ServiceType      string `xml:"device>serviceList>service>serviceType"`
	ServiceId        string `xml:"device>serviceList>service>serviceId"`
	DeviceType       string `xml:"device>deviceType"`
	ModelNumber      string `xml:"device>modelNumber"`
	URLBase          string `xml:"URLBase"`
	ModelURL         string `xml:"device>modelURL"`
	EventSubURL      string `xml:"device>serviceList>service>eventSubURL"`
	UDN              string `xml:"device>UDN"`
	PresentationURL  string `xml:"device>presentationURL"`
	ControlURL       string `xml:"device>serviceList>service>controlURL"`
	SCPDURL          string `xml:"device>serviceList>service>SCPDURL"`
	Manufacturer     string `xml:"device>manufacturer"`
	Xmlns            string `xml:"xmlns,attr"`
	Minor            string `xml:"specVersion>minor"`
	SerialNumber     string `xml:"device>serialNumber"`
	FriendlyName     string `xml:"device>friendlyName"`
}
