package log

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type loggerKey struct{}

// New ... new logger
func New(env string) (*zap.Logger, error) {
	if env == "development" {
		return newDevelopmentConfig().Build()
	}
	return newProductionConfig().Build()
}

// Logger ... get context from context
func logger(ctx context.Context) *zap.Logger {
	return ctxzap.Extract(ctx)
}

// Debugf ... output debug log
func Debugf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

// Infof ... output info log
func Infof(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

// Warningf ... output warning log
func Warningf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

// Errorf ... output error log
func Errorf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}
