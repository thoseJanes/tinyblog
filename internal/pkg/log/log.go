//go:generate
package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugw(msg string, keyAndValues ...interface{})
	Infow(msg string, keyAndValues ...interface{})
	Warnw(msg string, keyAndValues ...interface{})
	Errorw(msg string, keyAndValues ...interface{})
	Panicw(msg string, keyAndValues ...interface{})
	Fatalw(msg string, keyAndValues ...interface{})
	Sync()
	Clone() *Logger
}

type zapLogger struct{
	z *zap.Logger
}

var(
	mu sync.Mutex
	std = NewLogger(nil)
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *zapLogger {
	var encodingConfig = zap.NewProductionEncoderConfig()
	encodingConfig.MessageKey = "message"
	encodingConfig.TimeKey = "timestamp"
	encodingConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder){
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encodingConfig.EncodeDuration = func(t time.Duration, enc zapcore.PrimitiveArrayEncoder){
		enc.AppendFloat64(float64(t)/float64(time.Millisecond))
	}

	if opts == nil {
		opts = NewOptions()
	}
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(opts.LogLevel)); err!=nil {
		logLevel = zapcore.InfoLevel
	}

	config := zap.Config{
		DisableCaller: opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Encoding: opts.Format,
		EncoderConfig: encodingConfig,
		Level: zap.NewAtomicLevelAt(logLevel),

		OutputPaths: opts.OutputPaths,
		ErrorOutputPaths: opts.ErrOutputPaths,
	}
	logger, err := config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	zap.RedirectStdLog(logger)

	return &zapLogger{logger}
}

func Sync() {
	std.z.Sync()
}

func (l *zapLogger) Sync() {
	l.z.Sync()
}

func Clone() zapLogger {
	lc := *std
	return lc
}

func (l *zapLogger) Clone() zapLogger {
	lc := *l
	return lc
}

