package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type defaultLogger struct {
}

func NewDefaultLogger() InfDefine {
	return &defaultLogger{}
}

func (l *defaultLogger) CtxDebugf(ctx context.Context, format string, args ...interface{}) {
	logrus.Debugf(fmt.Sprintf("%+v %s", ctx, format), args...)
}

func (l *defaultLogger) CtxErrorf(ctx context.Context, format string, args ...interface{}) {
	logrus.Errorf(fmt.Sprintf("%+v %s", ctx, format), args...)
}

func (l *defaultLogger) CtxInfof(ctx context.Context, format string, args ...interface{}) {
	logrus.Infof(fmt.Sprintf("%+v %s", ctx, format), args...)
}

func (l *defaultLogger) CtxFatalf(ctx context.Context, format string, args ...interface{}) {
	logrus.Fatalf(fmt.Sprintf("%+v %s", ctx, format), args...)
}

func (l *defaultLogger) CtxTracef(ctx context.Context, format string, args ...interface{}) {
	logrus.Tracef(fmt.Sprintf("%+v %s", ctx, format), args...)
}
