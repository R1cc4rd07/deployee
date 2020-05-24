package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Logging struct
type Logging struct {
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "info":
		log.SetLevel(log.InfoLevel)
		break
	case "debug":
		log.SetLevel(log.DebugLevel)
		break
	case "error":
		log.SetLevel(log.ErrorLevel)
		break
	case "warn":
		log.SetLevel(log.WarnLevel)
		break
	case "fatal":
		log.SetLevel(log.FatalLevel)
		break
	case "panic":
		log.SetLevel(log.PanicLevel)
		break
	case "trace":
		log.SetLevel(log.TraceLevel)
		break
	default:
		log.SetLevel(log.InfoLevel)
		break
	}
	log.Printf("Loglevel: %s", log.GetLevel())
}

// Debugf Debugf function to print out a debug massage as a String and argurments as an interfacestruct
func (l *Logging) Debugf(s string, i ...interface{}) {
	log.Debugf(s, i)
}

// Debug function to print out a debug message gets argurments as an interface truct
func (l *Logging) Debug(i ...interface{}) {
	log.Debug(i)
}

// Errorf function to print out a error massage as a String and argurments as an interfacestruct
func (l *Logging) Errorf(s string, i ...interface{}) {
	log.Errorf(s, i)
}

// Error function to print out a error message gets argurments as an interface truct
func (l *Logging) Error(i ...interface{}) {
	log.Error(i)
}

// Infof function to print out a info massage as a String and argurments as an interface truct
func (l *Logging) Infof(s string, i ...interface{}) {
	log.Infof(s, i)
}

// Info function to print out a info message gets argurments as an interface sruct
func (l *Logging) Info(i ...interface{}) {
	log.Info(i)
}

// Fatalf function to print out a fatal massage as a String and argurments as an interfacestruct
func (l *Logging) Fatalf(s string, i ...interface{}) {
	log.Fatalf(s, i)
}

// Fatal function to print out a fatal message gets argurments as an interface truct
func (l *Logging) Fatal(i ...interface{}) {
	log.Fatal(i)
}

// Warnf function to print out a warn massage as a String and argurments as an interface truct
func (l *Logging) Warnf(s string, i ...interface{}) {
	log.Warnf(s, i)
}

// Warn function to print out a warn message gets argurments as an interface sruct
func (l *Logging) Warn(i ...interface{}) {
	log.Warn(i)
}

// Printf function to print out a print massage as a String and argurments as an interfacestruct
func (l *Logging) Printf(s string, i ...interface{}) {
	log.Printf(s, i)
}

// Print function to print out a print message gets argurments as an interface truct
func (l *Logging) Print(i ...interface{}) {
	log.Print(i)
}

// Println function to print out a print message gets argurments as an interfac struct
func (l *Logging) Println(i ...interface{}) {
	log.Println(i)
}
