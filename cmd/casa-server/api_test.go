package main

import (
	"casa-api/models"
	"casa-api/pkg/httpnet"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
	"time"
)

var baseURL = "http://127.0.0.1:3030/"

func startServer(t *testing.T) *exec.Cmd {
	os.Remove("./casa.db")
	os.Remove("./casa.db.lock")

	args := []string{"--port", "3030"}
	cmd := exec.Command("magnusson-home-security-server.exe", args...)
	err := cmd.Start()
	if err != nil {
		t.Fatal("Failed to start server", err)
	}

	return cmd
}

func TestAlarmsAndAlerts(t *testing.T) {
	alarmID := "hubba"
	alarmBaseURL := fmt.Sprintf("%sv1/alarms/", baseURL)
	alarmIDURL := fmt.Sprintf("%s%s", alarmBaseURL, alarmID)
	alarmToggleURL := fmt.Sprintf("%s/toggle", alarmIDURL)
	alarmLogsURL := fmt.Sprintf("%s/toggle", alarmIDURL)

	alertBaseURL := fmt.Sprintf("%sv1/alarms/%s/alerts/", baseURL, alarmID)
	alertAllBaseURL := fmt.Sprintf("%sv1/alerts/", baseURL)
	alertLivingRoomURL := fmt.Sprintf("%s%s", alertBaseURL, "livingroom")
	alertOfficeURL := fmt.Sprintf("%s%s", alertBaseURL, "office")

	cmd := startServer(t)
	defer cmd.Process.Kill()
	time.Sleep(time.Second * 2)

	devices := []string{"device1", "device2"}
	createAlarm(t, alarmBaseURL, alarmID, "create alarm for test", devices)
	assertCreatedAlarm(t, alarmIDURL, alarmID, "create alarm for test", devices)

	postAlarm(t, alarmToggleURL, true)
	postAlarm(t, alarmToggleURL, true)
	assertGetAlarm(t, alarmIDURL, alarmID, true)

	postAlarm(t, alarmToggleURL, false)
	assertGetAlarm(t, alarmIDURL, alarmID, false)

	assertListAlarms(t, alarmLogsURL, 2)

	// should not generate an alert when alarm is off
	postAlert(t, alertLivingRoomURL)
	postAlert(t, alertOfficeURL)

	assertListAlerts(t, alertLivingRoomURL, 0, 404)
	assertListAlerts(t, alertOfficeURL, 0, 404)

	// turn on alarm
	postAlarm(t, alarmToggleURL, true)
	assertGetAlarm(t, alarmIDURL, alarmID, true)
	assertListAlarms(t, alarmLogsURL, 3)

	// should generate an alert when alarm is on
	postAlert(t, alertOfficeURL)
	postAlert(t, alertLivingRoomURL)

	assertListAlerts(t, alertLivingRoomURL, 1, 200)
	assertListAlerts(t, alertOfficeURL, 1, 200)
	assertListAlerts(t, alertBaseURL, 2, 200)
	assertListAlerts(t, alertAllBaseURL, 2, 200)

	deletAlarm(t, alarmIDURL)
	assetDeletedAlarm(t, alarmIDURL)

	assertListAlerts(t, alertLivingRoomURL, 0, 404)
	assertListAlerts(t, alertOfficeURL, 0, 404)
}

func createAlarm(t *testing.T, url string, id string, description string, devices []string) {
	alarm := models.CreateAlarm{Description: description, ID: id, Devices: devices}
	resp, err := httpnet.Do(url, "POST", alarm, nil)
	if err != nil {
		t.Error("Create alarm failed with error:", err)
	}
	if resp.StatusCode != 201 {
		t.Fatalf("Create alarm failed. Status code: %d", resp.StatusCode)
	}
}

func postAlarm(t *testing.T, url string, on bool) {
	alarm := models.ToggleAlarm{On: on}
	resp, err := httpnet.Do(url, "POST", alarm, nil)
	if err != nil {
		t.Error("Toggle alarm failed with error:", err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("Toggle alarm failed. Status code: %d", resp.StatusCode)
	}
}

func postAlert(t *testing.T, url string) {
	alert := models.AddAlert{Description: "Desc", ShortDescription: "ShortDesc"}
	resp, err := httpnet.Do(url, "POST", alert, nil)
	if err != nil {
		t.Error("Add alert failed with error:", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Add alert failed. Status code: %d", resp.StatusCode)
	}
}

func deletAlarm(t *testing.T, url string) {
	resp, err := httpnet.Do(url, "DELETE", nil, nil)
	if err != nil {
		t.Error("Delete alarm with error:", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("Delete alarm failed. Status code: %d", resp.StatusCode)
	}
}

func assetDeletedAlarm(t *testing.T, url string) {
	var alarmsRes models.Alarm
	resp, err := httpnet.Do(url, "GET", nil, &alarmsRes)
	if err != nil {
		t.Error("Get alarm with error:", err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("The alarm was not delete. StatusCode: %d", resp.StatusCode)
	}
}

func assertGetAlarm(t *testing.T, url, expectedID string, expectedOn bool) {
	var alarmsRes models.Alarm
	resp, err := httpnet.Do(url, "GET", nil, &alarmsRes)
	if err != nil {
		t.Error("Failed to get alarm error:", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Get alarm failed. Status code: %d", resp.StatusCode)
	}
	if alarmsRes.ID != expectedID {
		t.Error("Alarm id should be hubba", alarmsRes.ID)
	}
	if alarmsRes.On != expectedOn {
		t.Errorf("Alarm should be %t but was %t", expectedOn, alarmsRes.On)
	}
	if len(alarmsRes.Date.String()) <= 0 {
		t.Error("Alarm should have a date")
	}
}

func assertCreatedAlarm(t *testing.T, url, expectedID string, expectedDescription string, expectedDevices []string) {
	var alarmsRes models.Alarm
	resp, err := httpnet.Do(url, "GET", nil, &alarmsRes)
	if err != nil {
		t.Error("Failed to get alarm error:", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Get alarm failed. Status code: %d", resp.StatusCode)
	}
	if alarmsRes.ID != expectedID {
		t.Errorf("Alarm id should be %s but was: %s", expectedID, alarmsRes.ID)
	}
	if alarmsRes.On != false {
		t.Errorf("Alarm should be %t but was %t", false, alarmsRes.On)
	}

	if alarmsRes.Description != expectedDescription {
		t.Errorf("Alarm should be %s but was %s", expectedDescription, alarmsRes.Description)
	}

	if !reflect.DeepEqual(expectedDevices, alarmsRes.Devices) {
		t.Errorf("Alarm devices should be %+v but was %+v", expectedDevices, alarmsRes.Devices)
	}

	if len(alarmsRes.Date.String()) <= 0 {
		t.Error("Alarm should have a date")
	}
}

func assertListAlerts(t *testing.T, url string, numberOfAlerts int, expectedStatusCode int) {
	var listAlertsRes []models.Alert
	resp, err := httpnet.Do(url, "GET", nil, &listAlertsRes)
	if err != nil {
		t.Error("Failed to get alerts error:", err)
	}
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Get alert failed. Status code: %d", resp.StatusCode)
	}
	if len(listAlertsRes) != numberOfAlerts {
		t.Errorf("Should have %d alerts but got %d", numberOfAlerts, len(listAlertsRes))
	}
}

func assertListAlarms(t *testing.T, url string, expectedCount int) {
	var listAlarmsRes []models.Alarm

	resp, err := httpnet.Do(url, "GET", nil, &listAlarmsRes)
	if err != nil {
		t.Error("List alarm failed:", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Toggle alarm failed. Status code: %d", resp.StatusCode)
	}

	if len(listAlarmsRes) != expectedCount {
		t.Errorf("Should have %d alarm logs but got %d", expectedCount, len(listAlarmsRes))
	}
}
