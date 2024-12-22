package logger

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	red   = "\033[31m" // Error
	green = "\033[32m" // Info
	blue  = "\033[34m" // Debug
	reset = "\033[0m"  // Reset color
)

// CustomLogger wraps zap.Logger and implements the Logger interface with default fields.
type CustomLogger struct {
	zapLogger     *zap.Logger
	defaultFields map[string]interface{}
}

// NewCustomLogger creates a new CustomLogger with optional default fields.
func NewCustomLogger(defaultFields map[string]interface{}) *CustomLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// Create a core with JSON encoding and output to stdout
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zapcore.InfoLevel,
	)

	// Create a logger with the core and add caller information
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &CustomLogger{
		zapLogger:     logger,
		defaultFields: defaultFields,
	}
}

// SetDefaultFields allows setting global fields that will be included in every log.
func (l *CustomLogger) SetDefaultFields(fields map[string]interface{}) {
	l.defaultFields = fields
}

// addFields converts a map to zap fields and merges them with default fields.
func (l *CustomLogger) addFields(extraFields ...map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0)

	for k, v := range l.defaultFields {
		fields = append(fields, zap.Any(k, v))
	}

	if len(extraFields) > 0 {
		for k, v := range extraFields[0] {
			fields = append(fields, zap.Any(k, v))
		}
	}

	return fields
}

// colorMessage applies color based on log level.
func (l *CustomLogger) colorMessage(level zapcore.Level, msg string) string {
	switch level {
	case zapcore.DebugLevel:
		return fmt.Sprintf("%s%s%s", blue, msg, reset)
	case zapcore.InfoLevel:
		return fmt.Sprintf("%s%s%s", green, msg, reset)
	case zapcore.ErrorLevel:
		return fmt.Sprintf("%s%s%s", red, msg, reset)
	default:
		return msg
	}
}

// Debug logs a message at the Debug level with optional extra fields.
func (l *CustomLogger) Debugc(ctx context.Context, msg string, extraFields map[string]interface{}) {
	extraFields["traceID"] = ctx.Value("tracectx")
	coloredMsg := l.colorMessage(zapcore.DebugLevel, msg)
	l.zapLogger.Debug(coloredMsg, l.addFields(extraFields)...)
}

// Info logs a message at the Info level with optional extra fields.
func (l *CustomLogger) Infoc(ctx context.Context, msg string, extraFields map[string]interface{}) {
	extraFields["traceID"] = ctx.Value("tracectx")
	coloredMsg := l.colorMessage(zapcore.InfoLevel, msg)
	l.zapLogger.Info(coloredMsg, l.addFields(extraFields)...)
}

// Warn logs a message at the Warn level with optional extra fields.
func (l *CustomLogger) Warnc(ctx context.Context, msg string, extraFields map[string]interface{}) {
	extraFields["traceID"] = ctx.Value("tracectx")
	l.zapLogger.Warn(msg, l.addFields(extraFields)...)
}

// Errorc logs a message at the Error level with optional extra fields.
func (l *CustomLogger) Errorc(ctx context.Context, msg string, extraFields map[string]interface{}) {
	extraFields["traceID"] = ctx.Value("tracectx")
	coloredMsg := l.colorMessage(zapcore.ErrorLevel, msg)
	l.zapLogger.Error(coloredMsg, l.addFields(extraFields)...)
}

// Sync flushes any buffered log entries.
func (l *CustomLogger) Sync() error {
	return l.zapLogger.Sync()
}
