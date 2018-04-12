package subscribers

import "casa/src/server/pkg/publishers"

func Enable() {
	motionDetectionSubscriber := motionDetectionHandler{}
	publishers.RegsiterAlarmEvent(motionDetectionSubscriber)

	smsSubscriber := smsSubscriber{}
	publishers.RegisterAlertEvent(smsSubscriber)
}
