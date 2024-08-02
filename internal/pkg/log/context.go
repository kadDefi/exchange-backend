package log

import (
	"context"

	"go.uber.org/zap"
)

type contextKey struct{}

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(contextKey{}).(*zap.Logger)
	if !ok {
		l = defaultLogger
	}

	return l
}
