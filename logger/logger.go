package logger

// Logger interface for the log package to easily change the logging bibliothek later on
type Logger interface {
	Debugf(s string, i ...interface{})
	Debug(i ...interface{})
	Warnf(s string, i ...interface{})
	Warn(i ...interface{})
	Infof(s string, i ...interface{})
	Info(i ...interface{})
	Fatalf(s string, i ...interface{})
	Fatal(i ...interface{})
	Errorf(s string, i ...interface{})
	Error(i ...interface{})
	Printf(s string, i ...interface{})
	Print(i ...interface{})
	Println(i ...interface{})
}
