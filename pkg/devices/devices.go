package devices

import (
	"casa-api/models"
	"casa-api/pkg/database"
	"casa-api/pkg/discovery"
	"casa-api/pkg/status"
	"casa-api/pkg/vapix"
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

var (
	devicesBucket = "devices"
	devicesKey    = "devicesMap"
)

func CreateBucketIfNotExists() error {
	err := database.CreateBucketIfNotExists(devicesBucket)
	if err != nil {
		return err
	}
	return nil
}

type device struct {
	SerialNumber string
	Username     string
	Password     string
}

type Device struct {
	SerialNumber   string
	Username       string
	Password       string
	Address        string
	ModelName      string
	ModelNumber    string
	Discovered     bool
	LastDiscovered time.Time
	Status         status.Status
}

func StartStatusRoutine(interval time.Duration) {
	go func() {
		for {
			devices, err := List()
			if err == nil {
				for _, device := range devices {
					status.UpdateStatus(device.SerialNumber, device.Address, device.Username, device.Password)
				}
			}
			time.Sleep(interval)
		}
	}()
}

func Save(addDevice *models.AddDevice) (int, error) {
	foundedDevice, ok := discovery.GetDiscoveredDevice(addDevice.SerialNumber)
	if ok {
		statusCode, err := vapix.CheckCredentials(foundedDevice.Address, addDevice.Credentials.Username, addDevice.Credentials.Password)
		if err != nil {
			logrus.Warnf("Failed to CheckCredentials for device with serialnumber: %s, err: %s", addDevice.SerialNumber, err.Error())
		}
		if statusCode == 200 {
			device := device{
				SerialNumber: addDevice.SerialNumber,
				Username:     addDevice.Credentials.Username,
				Password:     addDevice.Credentials.Password,
			}
			err = persist(device)
			if err != nil {
				return 500, errors.New("Failed to save the device")
			}
			return 200, nil
		}
		logrus.Warnf("Failed to CheckCredentials for device with serialnumber: %s, statusCode: %d", addDevice.SerialNumber, statusCode)
		return statusCode, errors.New("Failed to contact the device")
	}
	return 404, errors.New("Could not find the device")
}

func List() (map[string]Device, error) {
	devices := make(map[string]device)
	devicesToReturn := make(map[string]Device)
	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		v := b.Get([]byte(devicesKey))
		if len(v) > 0 {
			err := json.Unmarshal(v, &devices)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := database.View(f)
	if err != nil {
		return devicesToReturn, err
	}
	for _, d := range devices {
		deviceToReturn := Device{
			Password:     d.Password,
			Username:     d.Username,
			SerialNumber: d.SerialNumber,
		}
		discoveredDevice, found := discovery.GetDiscoveredDevice(d.SerialNumber)
		if found {
			deviceToReturn.Address = discoveredDevice.Address
			deviceToReturn.ModelName = discoveredDevice.ModelName
			deviceToReturn.ModelNumber = discoveredDevice.ModelNumber
			deviceToReturn.LastDiscovered = discoveredDevice.LastDiscovered
		}
		deviceToReturn.Discovered = found
		deviceToReturn.Status = status.GetStatus(deviceToReturn.SerialNumber)
		devicesToReturn[deviceToReturn.SerialNumber] = deviceToReturn
	}
	return devicesToReturn, nil
}

func Delete(serialNumber string) error {
	devices := make(map[string]device)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		v := b.Get([]byte(devicesKey))
		if len(v) > 0 {
			err := json.Unmarshal(v, &devices)
			if err != nil {
				return err
			}
			delete(devices, serialNumber)

			encoded, err := json.Marshal(devices)
			if err != nil {
				return err
			}

			return b.Put([]byte(devicesKey), encoded)
		}
		return nil
	}
	return database.Update(f)
}

func persist(deviceToPersist device) error {
	var dbUpdateFunc = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		devices := b.Get([]byte("devicesMap"))
		var existingDevices map[string]device
		if len(devices) > 0 {
			err := json.Unmarshal(devices, &existingDevices)
			if err != nil {
				return err
			}
		} else {
			existingDevices = make(map[string]device)
		}

		existingDevices[deviceToPersist.SerialNumber] = deviceToPersist

		encoded, err := json.Marshal(existingDevices)
		if err != nil {
			return err
		}

		return b.Put([]byte(devicesKey), encoded)
	}

	err := database.Update(dbUpdateFunc)
	if err != nil {
		logrus.Warnf("Failed to update device with serialnumber: %s, err: %s", deviceToPersist.SerialNumber, err.Error())
		return err
	}
	return nil
}
