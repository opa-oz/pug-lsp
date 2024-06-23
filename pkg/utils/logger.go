package utils

import (
	"io"
	"log"

	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/simple"
)

// Logger can be used to report logging messages.
// The native log.Logger type implements this interface.
type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Err(error)
}

type FileLogger struct {
	Logger *log.Logger
}

func NewFileLogger(file io.Writer, prefix string, flags int) *FileLogger {
	return &FileLogger{
		Logger: log.New(file, prefix, flags),
	}
}

func (l *FileLogger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

func (l *FileLogger) Println(v ...interface{}) {
	l.Logger.Println(v...)
}

func (l *FileLogger) Err(err error) {
	l.Logger.Printf("error: %v", err)
}

type SilentLogger struct {
	Logger *commonlog.Logger
}

func NewSilentLogger(prefix string) *SilentLogger {
	logger := commonlog.GetLogger(prefix)
	return &SilentLogger{
		Logger: &logger,
	}
}

func (l *SilentLogger) Printf(format string, v ...interface{}) {
	(*l.Logger).Noticef(format, v...)
}

func (l *SilentLogger) Println(v ...interface{}) {
	(*l.Logger).Notice("Println:", v...)
}

func (l *SilentLogger) Err(err error) {
	(*l.Logger).Errorf("error: %v", err)
}
