package subscribers

import "casa-api/pkg/publishers"

func Enable() {
	motionDetectionSubscriber := motionDetectionHandler{}
	publishers.RegsiterAlarmEvent(motionDetectionSubscriber)

	smsSubscriber := smsSubscriber{}
	publishers.RegisterAlertEvent(smsSubscriber)
}
