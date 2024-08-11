package logger

import (
	"fmt"
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Logger failed to init")
		return
	}
	Log = logger
}
