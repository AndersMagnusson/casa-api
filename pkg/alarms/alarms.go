package alarms

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"

	"casa/src/server/models"
	"casa/src/server/pkg/database"
	"casa/src/server/pkg/publishers"

	strfmt "github.com/go-openapi/strfmt"
)

var (
	alarmBucket  = "alarms"
	maxLogLength = 100
)

//var alarmLogBucket = "alarmslog"
// var currentLogKey = "current"
// var logsKey = "log"

func getLogsKey(alarmID string) string {
	return alarmID + "_logs"
}

func CreateBucketIfNotExists() error {
	err := database.CreateBucketIfNotExists(alarmBucket)
	if err != nil {
		return err
	}
	return nil
}

func GetAlarms() ([]models.Alarm, error) {
	alarms := make([]models.Alarm, 0)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alarmBucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if strings.Index(string(k), "_logs") < 0 {
				var alarm models.Alarm
				err := json.Unmarshal(v, &alarm)
				if err != nil {
					return err
				}
				alarms = append(alarms, alarm)
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	}
	err := database.View(f)
	return alarms, err
}

func GetAlarm(id string) (models.Alarm, error) {
	var alarm models.Alarm
	fmt.Println("GetAlarm: ", id)
	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alarmBucket))
		v := b.Get([]byte(id))
		if len(v) > 0 {
			err := json.Unmarshal(v, &alarm)
			if err != nil {
				return err
			}
		}
		fmt.Println("Alarm: ", alarm)
		return nil
	}
	err := database.View(f)
	return alarm, err
}

func DeleteAlarm(tx *bolt.Tx, id string) error {
	b := tx.Bucket([]byte(alarmBucket))
	err := b.Delete([]byte(id))
	if err != nil {
		return err
	}
	err = b.Delete([]byte(getLogsKey(id)))
	if err != nil {
		return err
	}
	return nil
}

func GetLogs(alarmID string) ([]models.ToggleAlarm, error) {
	alarms := make([]models.ToggleAlarm, 0)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alarmBucket))
		v := b.Get([]byte(getLogsKey(alarmID)))
		if len(v) > 0 {
			err := json.Unmarshal(v, &alarms)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := database.View(f)
	return alarms, err
}

func CreateAlarm(createAlarm *models.CreateAlarm) error {

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alarmBucket))
		v := b.Get([]byte(createAlarm.ID))
		var alarm models.Alarm

		if len(v) > 0 {
			err := json.Unmarshal(v, &alarm)
			if err != nil {
				return err
			}
		} else {
			alarm = models.Alarm{}
			alarm.ID = createAlarm.ID
		}
		alarm.Description = createAlarm.Description
		alarm.Devices = createAlarm.Devices
		alarm.MotionDetection = createAlarm.MotionDetection
		alarm.Continous = createAlarm.Continous
		alarm.Sms = createAlarm.Sms
		alarm.Date = strfmt.DateTime(time.Now())
		encoded, err := json.Marshal(alarm)
		if err != nil {
			return err
		}
		err = b.Put([]byte(alarm.ID), encoded)
		if err != nil {
			return err
		}
		return nil
	}
	return database.Update(f)
}

func ToggleAlarm(alarmID string, on bool) (*models.ToggleAlarm, error) {
	current, err := GetAlarm(alarmID)
	if err != nil {
		return nil, err
	}

	if current.On != on {
		toggleAlarm := models.ToggleAlarm{On: on, Date: strfmt.DateTime(time.Now()), AlarmID: alarmID}
		var dbUpdateFunc = func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(alarmBucket))
			current.On = on
			encoded, err := json.Marshal(current)
			if err != nil {
				return err
			}
			b.Put([]byte(alarmID), encoded)

			logs := b.Get([]byte(getLogsKey(alarmID)))
			var existingLogs []models.ToggleAlarm

			if len(logs) > 0 {
				err = json.Unmarshal(logs, &existingLogs)
				if err != nil {
					return err
				}
			} else {
				existingLogs = make([]models.ToggleAlarm, 0)
			}
			existingLogs = append(existingLogs, toggleAlarm)

			logLength := len(existingLogs)
			if logLength > maxLogLength {
				existingLogs = existingLogs[logLength-maxLogLength : logLength]
			}

			encoded, err = json.Marshal(existingLogs)
			if err != nil {
				return err
			}

			err = b.Put([]byte(getLogsKey(alarmID)), encoded)
			if err == nil {
				if on {
					publishers.On(current)
				}
				publishers.Off(current)
			}
			return err
		}

		return &toggleAlarm, database.Update(dbUpdateFunc)
	}
	return nil, nil
}
