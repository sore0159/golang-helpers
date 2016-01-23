package mylog

import (
	"fmt"
	"io"
	"mule/mybad"
	"os"
	"path"
	"runtime"
)

func (lg *Logger) Ping(v ...interface{}) {
	_, fName, lineNum, ok := runtime.Caller(1)
	var pingStr string
	if !ok {
		pingStr = "PING (?? ??)"
	} else {
		pingStr = fmt.Sprintf("PING (%s %02d)", path.Base(fName), lineNum)
	}
	lg.Println(append([]interface{}{pingStr}, v...)...)
}

func FilesWriter(fNames ...string) (writer io.Writer, err error) {
	files := make([]io.Writer, len(fNames))
	for i, filename := range fNames {
		xFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if my, bad := mybad.Check(err, "fileswriter failure", "filename", filename, "package", "mylog"); bad {
			return nil, my
		}
		files[i] = xFile
	}
	if len(files) == 1 {
		return files[0], nil
	}
	return io.MultiWriter(files...), nil
}

func (lg *Logger) AddFiles(fNames ...string) (*Logger, error) {
	w, err := FilesWriter(fNames...)
	if my, bad := mybad.Check(err, "logger addfiles failure"); bad {
		return nil, my
	}
	return lg.AddWriters(w), nil
}
func (lg *Logger) AddStderr() *Logger {
	return lg.AddWriters(os.Stderr)
}
func (lg *Logger) AddStdout() *Logger {
	return lg.AddWriters(os.Stdout)
}
