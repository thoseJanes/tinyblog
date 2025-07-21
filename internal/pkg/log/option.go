package log

import (
	"go.uber.org/zap/zapcore"
)

const(
	FormatConsole string = "console"
	FormatJson string = "json"
)


type Options struct{
	DisableCaller bool
	DisableStacktrace bool
	LogLevel string
	Format string
	OutputPaths []string
	ErrOutputPaths []string
}


func NewOptions() *Options {
	return &Options{
		DisableCaller: false,
		DisableStacktrace: false,
		LogLevel: zapcore.InfoLevel.String(),
		Format: FormatConsole,
		OutputPaths: []string{"stdout"},
		ErrOutputPaths: []string{"stderr"},
	}
}