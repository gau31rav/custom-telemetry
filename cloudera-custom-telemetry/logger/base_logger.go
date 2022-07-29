package logger

import (
	filename "github.com/keepeye/logrus-filename"
	"github.com/sirupsen/logrus"
)

type LogFormatter struct {
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}
	return formatter.Format(entry)
}

func formatLogger(logger *logrus.Logger) *logrus.Logger {
	logger.SetFormatter(&LogFormatter{})
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	logger.AddHook(filenameHook)
	logger.SetLevel(logrus.TraceLevel)
	return logger
}

// NewLogger instantiates a logrus logger with basic format.
func NewLogger(requestID string, actorCrn string) *logrus.Entry {
	logger := formatLogger(logrus.New())
	entry := logrus.NewEntry(logger).WithFields(logrus.Fields{
		"TenantID":  "1234-5678-9101",
		"RequestID": "1234-5678-9111",
	})
	return entry
}
