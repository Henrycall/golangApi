package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warning  *log.Logger
	err  *log.Logger
	writer  io.Writer
}

func Newlogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer,p ,log.Ldate |log.Ltime )

	return &Logger{
		debug : log.New(writer,"DEBUG : ", logger.Flags()),
		info : log.New(writer,"DEBUG : ", logger.Flags()),
		warning : log.New(writer,"DEBUG : ", logger.Flags()),
		err : log.New(writer,"DEBUG : ", logger.Flags()),
		writer : writer,
	}

}

// Create non-formatted logs

func (l * Logger) Debug(v ...interface{}){
	l.debug.Println(v ...)
}
func (l * Logger) Info(v ...interface{}){
	l.info.Println(v ...)
}
func (l * Logger) Warning(v ...interface{}){
	l.warning.Println(v ...)
}
func (l * Logger) Err(v ...interface{}){
	l.err.Println(v ...)
}

// create Formate Enabled logs

func (l * Logger) Debugf(format string, v ...interface{}){
	l.debug.Printf(format,v...)
}
func (l * Logger) Infof(format string, v ...interface{}){
	l.info.Printf(format,v...)
}
func (l * Logger) Warningf(format string, v ...interface{}){
	l.warning.Printf(format,v...)
}
func (l * Logger) Errf(format string, v ...interface{}){
	l.err.Printf(format,v...)
}
