package adapters

import (
	"github.com/sirupsen/logrus"
)

type BasicLoggerAdapter struct {
}

// NewBasicLoggerAdapter - Create new BasicLoggerAdapter instance
func NewBasicLoggerAdapter() BasicLoggerAdapter {
	return BasicLoggerAdapter{}
}

// Info - Log data with Info logging level
func (b *BasicLoggerAdapter) Info(i interface{}) {
	logrus.Info(i)
}

// Warn - Log data with Warning logging level
func (b *BasicLoggerAdapter) Warn(i interface{}) {
	logrus.Warn(i)
}

// Error - Log data with Error logging level
func (b *BasicLoggerAdapter) Error(i interface{}) {
	logrus.Error(i)
}

// Debug - Log data with Debug logging level
func (b *BasicLoggerAdapter) Debug(i interface{}) {
	logrus.Debug(i)
}

// init - To configure application log
func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		DisableQuote:  true,
	})

	logrus.SetReportCaller(false)
	logrus.SetLevel(4)
}
