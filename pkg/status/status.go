package status

import (
	"casa-api/pkg/vapix"
	"sync"
	"time"
)

var deviceStatusesLock sync.Mutex
var deviceStatuses = make(map[string]Status)

type Status struct {
	Network         bool
	Credential      bool
	Error           bool
	LastStatusCheck time.Time
	HasStatus       bool
	StatusCode      int
	Message         string
	Username        string
	Password        string
}

func CheckStatus(serialNumber, address, username, password string) {
	status := check(address, username, password)

	deviceStatusesLock.Lock()
	deviceStatuses[serialNumber] = status
	deviceStatusesLock.Unlock()

}

func GetStatus(serialNumber string) Status {
	var temp Status
	deviceStatusesLock.Lock()
	status, ok := deviceStatuses[serialNumber]
	if ok {
		temp = status
	}
	deviceStatusesLock.Unlock()
	return temp
}

func UpdateDiscoveryStatus(serialNumber, address string) {
	deviceStatusesLock.Lock()
	existing, ok := deviceStatuses[serialNumber]
	deviceStatusesLock.Unlock()
	if ok {
		if existing.Credential {
			return
		}
	}
	UpdateStatus(serialNumber, address, existing.Username, existing.Password)
}

func UpdateStatus(serialNumber, address, username, password string) Status {
	status := check(address, username, password)
	deviceStatusesLock.Lock()
	deviceStatuses[serialNumber] = status
	deviceStatusesLock.Unlock()
	return status
}

func check(address, username, password string) Status {
	status := Status{
		Network:         false,
		Credential:      false,
		Error:           false,
		LastStatusCheck: time.Now(),
		HasStatus:       false,
		Message:         "unkown",
		Username:        username,
		Password:        password,
	}

	if len(address) <= 0 {
		return status
	}
	statusCode, _ := vapix.CheckCredentials(address, username, password)
	status.StatusCode = statusCode
	if statusCode == 200 {
		status.Network = true
		status.Credential = true
		status.Message = "ok"
		status.HasStatus = true
		status.Error = false
	} else if statusCode == 401 {
		status.Network = true
		status.Credential = false
		status.Message = "wrong credentials"
		status.HasStatus = true
		status.Error = false
	} else if statusCode > 0 {
		status.Network = true
		status.Credential = false
		status.Message = "unknown error"
		status.HasStatus = true
		status.Error = true
	}
	return status
}
