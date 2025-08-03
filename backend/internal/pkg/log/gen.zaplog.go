package log

var _ Logger = &zapLogger{}


func Debugw (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Debugw(msg, keyAndValues...)
}

func (l *zapLogger) Debugw (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Debugw(msg, keyAndValues...)
}

func Infow (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Infow(msg, keyAndValues...)
}

func (l *zapLogger) Infow (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Infow(msg, keyAndValues...)
}

func Warnw (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Warnw(msg, keyAndValues...)
}

func (l *zapLogger) Warnw (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Warnw(msg, keyAndValues...)
}

func Errorw (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Errorw(msg, keyAndValues...)
}

func (l *zapLogger) Errorw (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Errorw(msg, keyAndValues...)
}

func Panicw (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Panicw(msg, keyAndValues...)
}

func (l *zapLogger) Panicw (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Panicw(msg, keyAndValues...)
}

func Fatalw (msg string, keyAndValues ...interface{}) {
  std.z.Sugar().Fatalw(msg, keyAndValues...)
}

func (l *zapLogger) Fatalw (msg string, keyAndValues ...interface{}){
  l.z.Sugar().Fatalw(msg, keyAndValues...)
}

