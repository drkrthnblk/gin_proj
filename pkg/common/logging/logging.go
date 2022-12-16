package logging

import (
	"github.com/siruspen/logrus"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func SetFormat(format string) {
	switch format {
	case "json":
		setupJsonLogging()
	default:
		setupTextLogging()
	}
}

func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(lvl)
}

func setupJsonLogging() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFunc: "func",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg: "message",
			logrus.FieldKeyLogrusError: "error",
			logrus.FieldKeyTime: "time",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			arr := strings.Split(frame.Func.Name(), ".")
			return arr[len(arr)-1], fileName
		},
	})
	logrus.SetReportCaller(true)
}

func setupTextLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFunc: "func",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg: "message",
			logrus.FieldKeyLogrusError: "error",
			logrus.FieldKeyTime: "time",
		},
		PadLevelText: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
}