package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func flatten(fields map[string]interface{}) []interface{} {
	result := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		result = append(result, k, v)
	}
	return result
}

type ZapLogger struct {
	z *zap.SugaredLogger
}

func NewZapLogger() Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	base, _ := cfg.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	return &ZapLogger{z: base.Sugar()}
}

func (l *ZapLogger) Info(msg string, fields Fields) {
	kv := flatten(fields)
	l.z.Infow(msg, kv...)
}

func (l *ZapLogger) Debug(msg string, fields Fields) {
	kv := flatten(fields)
	l.z.Debugw(msg, kv...)
}

func (l *ZapLogger) Error(msg string, fields Fields) {
	kv := flatten(fields)
	l.z.Errorw(msg, kv...)
}

func (l *ZapLogger) Warn(msg string, fields Fields) {
	kv := flatten(fields)
	l.z.Warnw(msg, kv...)
}

func (l *ZapLogger) Fatal(msg string, fields Fields) {
	kv := flatten(fields)
	l.z.Fatalw(msg, kv...)
}
