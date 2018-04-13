package subscribers

import (
	"casa-api/models"
	"casa-api/pkg/httpnet"
	"fmt"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type smsSubscriber struct {
}

var lastSms = time.Time{}

func (h smsSubscriber) Alert(alarm models.Alarm, alert models.Alert) {
	if alarm.Sms != nil && len(alarm.Sms.MobileNumbers) > 0 {
		diff := time.Now().Sub(lastSms)
		logrus.Infof("SMS diff:", diff.Nanoseconds)
		if diff.Nanoseconds() > int64(alarm.Sms.Interval)*60*1000*1000*1000 {
			for _, toNumber := range alarm.Sms.MobileNumbers {
				smsURL := fmt.Sprintf("https://se-1.cellsynt.net/sms.php?username=%s&password=%s&destination=%s&originatortype=numeric&originator=%s&charset=UTF-8&text=%s",
					alarm.Sms.Username,
					alarm.Sms.Password,
					toNumber,
					alarm.Sms.FromNumber,
					url.QueryEscape(alert.ShortDescription))

				resp, err := httpnet.Do(smsURL, "GET", nil, nil)
				if err != nil {
					if err != nil {
						logrus.Warnf("Failed send sms, err: %s", err.Error())
						return
					}
				} else {
					if resp.StatusCode == 200 {
						logrus.Infof("Sent sms for alarm: %s, alert: %s", alarm.ID, alert.ID)
					} else {
						logrus.Warnf("Failed send sms, statusCode: %d, body: %s", resp.StatusCode, resp.Body)
					}
				}
			}
			lastSms = time.Now()
		}
	}
}
