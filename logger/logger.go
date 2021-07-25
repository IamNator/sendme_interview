package logger

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

//Logger logs information
var Logger *log.Logger
var Warn *log.Logger
var Fatal *log.Logger

// var logglyToken = ""

// var url = "http://logs-01.loggly.com/inputs//tag/http/"

func init() {

	Logger = log.New()

	Logger.SetFormatter(&log.JSONFormatter{PrettyPrint: false})

	Logger.WithFields(logrus.Fields{
		"appname": "sendme_interview",
	})

	Warn = log.New()
	Warn.Level = logrus.WarnLevel
	Warn.WithFields(logrus.Fields{
		"appname": "sendme_interview",
	})

	Fatal = log.New()
	Fatal.Level = logrus.FatalLevel
	Fatal.WithFields(logrus.Fields{
		"appname": "sendme_interview",
	})

	//log.SetOutput(OpenLogfile())

	Logger.Println("logger object created")
}

//InvalidArgValue logs errors in the following format
// -->  Invalid value for argument: %s: %v\n
func InvalidArgValue(client string, err interface{}) {
	Logger.Printf("Invalid value for argument: %s: %v\n", client, err)
}
