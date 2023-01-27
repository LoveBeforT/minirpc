package minirpc

import (
	"sync"

	"github.com/KouKouChan/minirpc/logger"
	"github.com/sirupsen/logrus"
)

var (
	loggerLock sync.Mutex
	loggerImpl logger.InfDefine
)

func init() {
	loggerLock.Lock()
	defer loggerLock.Unlock()
	loggerImpl = logger.NewDefaultLogger()
}

func SetLogger(logger logger.InfDefine) {
	loggerLock.Lock()
	defer loggerLock.Unlock()
	loggerImpl = logger
}

func SetLoggerLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func GetLogger() logger.InfDefine {
	loggerLock.Lock()
	defer loggerLock.Unlock()
	return loggerImpl
}
