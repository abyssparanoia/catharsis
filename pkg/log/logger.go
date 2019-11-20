package log

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type loggerKey struct{}

// New ... new logger
func New() (*zap.Logger, error) {
	return newProductionConfig().Build()
}

// NewDevelopment ... new development logger
func NewDevelopment() (*zap.Logger, error) {
	return newDevelopmentConfig().Build()
}

// Logger ... get context from context
func logger(ctx context.Context) *zap.Logger {
	return ctxzap.Extract(ctx)
}

// Debugf ... output debug log
func Debugf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).Debug(msg, fields...)
}

// Infof ... output info log
func Infof(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).Info(msg, fields...)
}

// Warningf ... output warning log
func Warningf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).Warn(msg, fields...)
}

// Errorf ... output error log
func Errorf(ctx context.Context, msg string, fields ...zap.Field) {
	logger(ctx).Error(msg, fields...)
}
