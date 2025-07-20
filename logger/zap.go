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

func NewZapLogger(withCaller bool) Logger {
	cfg := zap.NewDevelopmentConfig()

	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var options []zap.Option
	if withCaller {
		options = append(options,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		)
	} else {
		cfg.DisableCaller = true
	}

	base, _ := cfg.Build(options...)

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
