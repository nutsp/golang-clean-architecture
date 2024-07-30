package observability

import (
	"time"

	"github.com/nutsp/golang-clean-architecture/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Development string = "development"
	Production  string = "production"
)

type (
	Logger interface {
		Info(msg string, fields ...interface{})
		Error(msg string, fields ...interface{})
		Debug(msg string, fields ...interface{})
	}

	ZapLogger struct {
		logger *zap.SugaredLogger
	}
)

func NewZapLogger(cfg config.Logger) *ZapLogger {
	var config zap.Config

	if cfg.Mode == Development {
		config = zap.Config{
			Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
			Development: true,
			Encoding:    "console",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Capital level with color
				EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 UTC time format
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
	} else {
		config = zap.Config{
			Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Development: false,
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,             // Lowercase log level
				EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339), // Custom time format
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
	}

	zapLogger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync() // flushes buffer, if any

	sugaredLogger := zapLogger.Sugar()
	return &ZapLogger{
		logger: sugaredLogger,
	}
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Infow(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Errorw(msg, fields...)
}

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Debugw(msg, fields...)
}
