package publishers

import (
	"casa/src/server/models"
)

var alarmEvents = make([]AlarmEvent, 0)
var alertEvents = make([]AlertEvent, 0)

type AlarmEvent interface {
	On(alarm models.Alarm)
	Off(alarm models.Alarm)
}

type AlertEvent interface {
	Alert(alarm models.Alarm, alert models.Alert)
}

func RegsiterAlarmEvent(alarmEvent AlarmEvent) {
	alarmEvents = append(alarmEvents, alarmEvent)
}

func RegisterAlertEvent(alertEvent AlertEvent) {
	alertEvents = append(alertEvents, alertEvent)
}

func On(alarm models.Alarm) {
	for _, event := range alarmEvents {
		go event.On(alarm)
	}
}

func Off(alarm models.Alarm) {
	for _, event := range alarmEvents {
		go event.Off(alarm)
	}
}

func Alert(alarm models.Alarm, alert models.Alert) {
	for _, event := range alertEvents {
		go event.Alert(alarm, alert)
	}
}
