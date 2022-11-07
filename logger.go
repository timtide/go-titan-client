package titan_client

import (
	logging "github.com/ipfs/go-log/v2"
	"titan-client/common"
)

var logger = NewLog().getLogger()

type log struct {
	systemName string
}

func NewLog() *log {
	return &log{systemName: common.AppName}
}

// SetLevel set log level, eg: DEBUG, INFO, WARN, ERROR...
func (l *log) SetLevel(level string) error {
	return logging.SetLogLevel(l.systemName, level)
}

func (l *log) getLogger() *logging.ZapEventLogger {
	return logging.Logger(l.systemName)
}
