package alarms

import (
	"casa/src/server/models"
	"casa/src/server/pkg/database"
	"os"
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func setUp(t *testing.T) *bolt.DB {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		t.Fatal("Failed with bucket", err)
	}
	database.SetDatabase(db)
	err = CreateBucketIfNotExists()
	if err != nil {
		t.Fatal("Failed with bucket", err)
	}
	return db
}
func TestAdd(t *testing.T) {
	id := "test"
	db := setUp(t)
	defer tearDown(t, db)

	devices := []string{"1", "2"}
	createAlarm := models.CreateAlarm{ID: id, Description: "test", Devices: devices}
	err := CreateAlarm(&createAlarm)
	if err != nil {
		t.Error("Failed to create alarm", err)
	}

	existingAlarm, err := GetAlarm(id)
	if existingAlarm.Description != "test" {
		t.Error("Description was wrong", existingAlarm.Description)
	}

	if !reflect.DeepEqual(existingAlarm.Devices, devices) {
		t.Error("Devices were not correct", existingAlarm.Devices)
	}

	_, err = ToggleAlarm(id, true)
	if err != nil {
		t.Error("Failed to set alarm", err)
	}

	actual, err := GetAlarm(id)
	if err != nil {
		t.Error("Failed to get alarm", err)
	}

	if !actual.On {
		t.Error("Alarm should be on")
	}
}

func TestGetLogs(t *testing.T) {
	id := "test"
	db := setUp(t)
	defer tearDown(t, db)
	for index := 0; index < 10; index++ {
		err, _ := ToggleAlarm(id, true)
		if err != nil {
			t.Error("Failed to set alarm", err)
		}
	}

	actualLogs, err := GetLogs(id)
	if err != nil {
		t.Error("Failed to get alarm", err)
	}

	if len(actualLogs) != 10 {
		t.Error("Should have 10 logs")
	}
}

func tearDown(t *testing.T, db *bolt.DB) {
	err := db.Close()
	if err != nil {
		t.Error("Failed to close test db", err)
	}
	err = os.Remove("test.db")
	if err != nil {
		t.Error("Failed to remove test db", err)
	}
}
