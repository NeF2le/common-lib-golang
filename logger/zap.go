package logger

import "go.uber.org/zap"

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

func NewZapLogger(z *zap.Logger) Logger {
	return &ZapLogger{z: z.Sugar()}
}

func (l *ZapLogger) Info(msg string, fields Fields) {
	l.z.Infow(msg, flatten(fields))
}

func (l *ZapLogger) Debug(msg string, fields Fields) {
	l.z.Debugw(msg, flatten(fields))
}

func (l *ZapLogger) Error(msg string, fields Fields) {
	l.z.Errorw(msg, flatten(fields))
}

func (l *ZapLogger) Warn(msg string, fields Fields) {
	l.z.Warnw(msg, flatten(fields))
}

func (l *ZapLogger) Fatal(msg string, fields Fields) {
	l.z.Fatalw(msg, flatten(fields))
}
