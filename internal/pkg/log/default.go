package log

import (
	"os"

	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

func init() {
	switch os.Getenv("DEPLOYMENT_ENVIRONMENT") {
	case "production":
		{
			l, err := zap.NewProduction()
			if err != nil {
				panic(err)
			}
			defaultLogger = l

		}
	default:
		{
			l, err := zap.NewDevelopment()
			if err != nil {
				panic(err)
			}
			defaultLogger = l
		}
	}
}

func DefaultLogger() *zap.Logger {
	return defaultLogger
}
