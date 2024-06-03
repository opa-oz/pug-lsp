package utils

import (
	"io"
	"log"
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
