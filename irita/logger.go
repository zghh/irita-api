package irita

import (
	"irita-api/seelog"

	"github.com/tendermint/tendermint/libs/log"
)

type Logger struct {
}

func (l *Logger) Debug(msg string, keyvals ...interface{}) {
	seelog.Debugf(msg)
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	seelog.Infof(msg)
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	seelog.Errorf(msg)
}

func (l *Logger) With(keyvals ...interface{}) log.Logger {
	return l
}
