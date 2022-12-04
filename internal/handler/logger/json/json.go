package json

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
)

var _ logger.Logger = (*jsonLogger)(nil)

func NewLogger() (logger.Logger, error) {
	l, err := newZapLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to create json logger: %w", err)
	}

	return &jsonLogger{
		logger: zapr.NewLogger(l),
	}, nil
}

type jsonLogger struct {
	logger logr.Logger
}

func (l *jsonLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l *jsonLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(err, msg, keysAndValues...)
}

func (l *jsonLogger) WithValues(keysAndValues ...interface{}) logger.Logger {
	return &jsonLogger{
		logger: l.logger.WithValues(keysAndValues...),
	}
}

func (l *jsonLogger) WithName(name string) logger.Logger {
	return &jsonLogger{
		logger: l.logger.WithName(name),
	}
}

func newZapLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		Encoding:          "json",
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "severity",
			NameKey:        "logger",
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochMillisTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	l, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create zap logger: %w", err)
	}

	return l, nil
}
