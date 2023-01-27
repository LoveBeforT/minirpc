package logger

import "context"

type InfDefine interface {
	CtxDebugf(ctx context.Context, format string, args ...interface{})
	CtxErrorf(ctx context.Context, format string, args ...interface{})
	CtxInfof(ctx context.Context, format string, args ...interface{})
	CtxFatalf(ctx context.Context, format string, args ...interface{})
	CtxTracef(ctx context.Context, format string, args ...interface{})
}
