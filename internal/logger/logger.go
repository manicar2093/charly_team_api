package logger

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func Warning(content interface{}) {
	logrus.Warn(toJson(content))
}

func Info(content interface{}) {
	logrus.Info(toJson(content))
}

func Error(content interface{}) {
	logrus.Error(content)
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func toJson(d interface{}) string {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(b)
}
