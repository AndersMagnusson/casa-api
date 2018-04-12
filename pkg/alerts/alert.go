package alerts

import (
	"casa/src/server/models"
	"casa/src/server/pkg/alarms"
	"casa/src/server/pkg/database"
	"casa/src/server/pkg/publishers"
	"fmt"

	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/boltdb/bolt"
)

var (
	alertBucket  = "alerts"
	maxLogLength = 100
)

func CreateBucketIfNotExists() error {
	err := database.CreateBucketIfNotExists(alertBucket)
	if err != nil {
		return err
	}
	return nil
}

func HandleAlert(alarmID string, id string, addAlert models.AddAlert) error {
	alarm, err := alarms.GetAlarm(alarmID)
	if err != nil {
		return err
	}
	if alarm.On {
		alert := models.Alert{
			AlarmID:          alarmID,
			Description:      addAlert.Description,
			ShortDescription: addAlert.ShortDescription,
			Date:             strfmt.DateTime(time.Now()),
		}
		alert.ID = id

		var dbUpdateFunc = func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(alertBucket))
			logs := b.Get([]byte(alarmID))
			var existingLogs map[string][]models.Alert
			if len(logs) > 0 {
				err = json.Unmarshal(logs, &existingLogs)
				if err != nil {
					return err
				}
			} else {
				existingLogs = make(map[string][]models.Alert)
			}
			arrAlerts, ok := existingLogs[id]
			if ok {
				existingLogs[id] = append(arrAlerts, alert)
			} else {
				existingLogs[id] = []models.Alert{alert}
			}

			logLength := len(existingLogs)
			if logLength > maxLogLength {
				existingLogs[id] = existingLogs[id][logLength-maxLogLength : logLength]
			}

			encoded, err := json.Marshal(existingLogs)
			if err != nil {
				return err
			}

			return b.Put([]byte(alarmID), encoded)

			// if len(existingLogs) > 0 {
			// 	// last := existingLogs[id][len(existingLogs)-1]
			// 	// diff := time.Now().Sub(time.Time(last.Date)).Nanoseconds()
			// 	// if diff >= 60*60*1000*1000*1000 {
			// 	encoded, err := json.Marshal(existingLogs)
			// 	if err != nil {
			// 		return err
			// 	}

			// 	return b.Put([]byte(alarmID), encoded)
			// 	// }
			// }
			// return nil
		}

		err := database.Update(dbUpdateFunc)
		if err != nil {
			logrus.Warnf("Failed to update alert log, err: %s", err.Error())
		}

		publishers.Alert(alarm, alert)

		return nil
		// TODO notify
	}
	return nil
}

func GetLogs(alarmID string, id string) ([]models.Alert, error) {
	alerts := make([]models.Alert, 0)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alertBucket))
		v := b.Get([]byte(alarmID))
		if len(v) > 0 {
			var alarmDidc map[string][]models.Alert
			err := json.Unmarshal(v, &alarmDidc)
			if err != nil {
				return err
			}
			existingAlertArr, ok := alarmDidc[id]
			if ok {
				alerts = existingAlertArr
			}
		}
		return nil
	}
	err := database.View(f)
	return alerts, err
}

func GetAllLogs() ([]models.Alert, error) {

	result := make([]models.Alert, 0)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alertBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			if len(v) > 0 {
				var alarmDidc map[string][]models.Alert
				err := json.Unmarshal(v, &alarmDidc)
				if err != nil {
					return err
				}
				for _, v := range alarmDidc {
					result = append(result, v...)
				}

			}
		}
		return nil
	}
	err := database.View(f)
	return result, err
}

func GetAllLogsForAlarm(alarmID string) ([]models.Alert, error) {

	result := make([]models.Alert, 0)

	var f = func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(alertBucket))
		v := b.Get([]byte(alarmID))
		if len(v) > 0 {
			var alarmDidc map[string][]models.Alert
			err := json.Unmarshal(v, &alarmDidc)
			if err != nil {
				return err
			}

			for _, v := range alarmDidc {
				result = append(result, v...)
			}
			// sort.Slice(result, func(i, j int) bool {
			// 	return result[i].Date.String() < result[j].Date.String()
			// })
		}
		return nil
		// c := b.Cursor()

		// for k, v := c.First(); k != nil; k, v = c.Next() {
		// 	fmt.Printf("key=%s, value=%s\n", k, v)
		// 	err := json.Unmarshal(v, &alerts)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if len(alerts) > 0 {
		// 		result = append(result, alerts...)
		// 	}
		// }
		// return nil
	}
	err := database.View(f)
	return result, err
}

func Delete(tx *bolt.Tx, alarmID string) error {
	b := tx.Bucket([]byte(alertBucket))
	err := b.Delete([]byte(alarmID))
	if err != nil {
		return err
	}
	return nil
}
