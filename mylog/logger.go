package mylog

import (
	"fmt"
	"io"
	"log"
)

// Logger provides an immutable alternative to
// the standard logger: changing outputs and
// settings simply spurs new *Loggers
//
// Loggers behave otherwise just like standard
// library loggers with an extra check on
// Println and Print for *MuleError handling
type Logger struct {
	*log.Logger
	writer io.Writer
}

func NewLogger() *Logger {
	return NULLLOG
}

func New(w io.Writer, pref string, flags int) *Logger {
	return &Logger{log.New(w, pref, flags), w}
}

func (lg *Logger) AddWriters(ws ...io.Writer) *Logger {
	var multi io.Writer
	if lg.writer != DEVNULL {
		ws = append(ws, lg.writer)
	}
	if len(ws) == 1 && lg.writer == DEVNULL {
		multi = ws[0]
	} else {
		multi = io.MultiWriter(ws...)
	}
	pref, flags := lg.Prefix(), lg.Flags()
	return &Logger{log.New(multi, pref, flags), multi}
}

func (lg *Logger) SetOutput(w io.Writer) *Logger {
	pref, flags := lg.Prefix(), lg.Flags()
	return &Logger{log.New(w, pref, flags), w}
}
func (lg *Logger) SetPrefix(pref string) *Logger {
	flags := lg.Flags()
	return &Logger{log.New(lg.writer, pref, flags), lg.writer}
}
func (lg *Logger) SetFlags(flags int) *Logger {
	pref := lg.Prefix()
	return &Logger{log.New(lg.writer, pref, flags), lg.writer}
}
func (lg *Logger) AddPrefix(pref2 string) *Logger {
	pref, flags := lg.Prefix(), lg.Flags()
	pref = fmt.Sprintf("%s%s", pref, pref2)
	return &Logger{log.New(lg.writer, pref, flags), lg.writer}
}
