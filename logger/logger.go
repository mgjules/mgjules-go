package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var L zap.SugaredLogger

func Init(prod bool) error {
	var (
		logger *zap.Logger
		err    error
	)
	if prod {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	L = *logger.Sugar()

	return nil
}
