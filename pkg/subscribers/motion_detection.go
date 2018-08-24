package subscribers

import (
	"casa-api/models"
	"casa-api/pkg/devices"
	"casa-api/pkg/status"
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
			logrus.Infof("Handling motion detection for %s", serialNumber)
			currentStatus := status.GetStatus(serialNumber)
			if currentStatus.IsOk() {
				motionDetectionHandler := axisevents.NewMotionDetectionHandler("Magnusson", on, currentStatus.VmdVersion)

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
						err = motionDetectionHandler.Record().Video().SdCard().ExecuteOn(context.TODO(), device)
						if err != nil {
							logrus.Warnf("Failed to toggle motion detection %t, err: %s", on, err.Error())
						} else {
							logrus.Infof("Motion detection on %s, %s (vmd version %d) was turned: %t", device.Address, serialNumber, currentStatus.VmdVersion, on)
						}
					}
				}
			} else {
				logrus.Warnf("Device is %s is not reachable and hence not able to toggle motion recording", serialNumber)
			}
		}
	}
}
