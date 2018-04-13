package subscribers

import (
	"casa-api/models"
	"casa-api/pkg/devices"
	"context"

	"github.com/AndersMagnusson/axisevents"
	"github.com/sirupsen/logrus"
)

type motionDetectionHandler struct {
}

func (h motionDetectionHandler) On(alarm models.Alarm) {
	h.handleEvent(true, alarm)
}

func (h motionDetectionHandler) Off(alarm models.Alarm) {
	h.handleEvent(false, alarm)
}

func (h motionDetectionHandler) handleEvent(on bool, alarm models.Alarm) {
	if alarm.MotionDetection {
		devices, err := devices.List()
		for _, serialNumber := range alarm.Devices {

			motionDetectionHandler := axisevents.NewMotionDetectionHandler("Magnusson", on)

			if err != nil {
				logrus.Warnf("Failed to set motion detection %t, err: %s", on, err.Error())
				return
			}

			for _, d := range devices {
				if d.SerialNumber == serialNumber && d.Discovered {
					device := axisevents.Device{
						Username: d.Username,
						Password: d.Password,
						Address:  d.Address,
					}
					motionDetectionHandler.Record().Video().SdCard().ExecuteOn(context.TODO(), device)
				}
			}
		}
	}
}
