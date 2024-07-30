package observability

import (
	"os"

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

	var encoder zapcore.Encoder
	var config zapcore.EncoderConfig

	if cfg.Mode == Development {
		config = zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	} else {
		config = zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	}

	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(config)
	} else {
		encoder = zapcore.NewJSONEncoder(config)
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zap.NewAtomicLevelAt(zapcore.DebugLevel))
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

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
