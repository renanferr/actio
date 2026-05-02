package logger

import "context"

type contextKey string

const loggerContextKey contextKey = "actio-logger"

func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return &SimpleLogger{}
	}
	if value := ctx.Value(loggerContextKey); value != nil {
		if logger, ok := value.(Logger); ok {
			return logger
		}
	}
	return &SimpleLogger{}
}
