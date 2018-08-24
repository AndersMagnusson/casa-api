package logging

import (
	"os"

	"github.com/jonathan-robertson/logentrus"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{}) // logentrus hook will always submit JSON to Logentries

	token := os.Getenv("CasaLogKey")

	hook, err := logentrus.New(
		token, // grabbing Logentries Token from environment variable
		&logentrus.Opts{
			Priority:        logrus.InfoLevel, // since set to InfoLevel, DebugLevel is the only level that will be ignored
			TimestampFormat: "",               // Jan 2 15:04:05 setting empty string here will default to logrus's typically time format
			EncTLSConfig:    nil,              // setting config to nil means that conn will use root certs from local system
			UnencryptedTCP:  false,            // disable encryption, but still use TCP
			UnencryptedUDP:  false,            // disable encryption and use UDP
			UnencryptedPort: 514,              // omitting will result in port 514 usage; valid options are 80, 514, and 10000
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(hook)
}
