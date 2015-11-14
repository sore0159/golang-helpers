package mylog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	defaultLogName    = "mainlogs.txt"
	defaultErrLogName = "errorlogs.txt"
)

var (
	MainLog *log.Logger
	ErrLog  *log.Logger
)

func InitDefaults() {
	SetErr(defaultErrLogName)
	SetMain(defaultLogName)
}

func ErrF(t string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(t, args...))
}

func Log(v ...interface{}) error {
	if len(v) == 1 {
		if err, ok := v[0].(error); ok {
			MainLog.Println(err)
			return err
		}
	}
	MainLog.Println(v...)
	return errors.New(fmt.Sprint(v...))
}

func Err(v ...interface{}) error {
	if len(v) == 1 {
		if err, ok := v[0].(error); ok {
			ErrLog.Println(err)
			return err
		}
	}
	ErrLog.Println(v...)
	return errors.New(fmt.Sprint(v...))
}

func SetErr(fileName string) error {
	mEF, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println("Failed to set Err:", err)
		ErrLog = log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime)
		return err
	}
	errMulti := io.MultiWriter(mEF, os.Stderr)
	ErrLog = log.New(errMulti, "ERROR:", log.Ldate|log.Ltime)
	return nil
}

func SetMain(fileName string) error {
	mLF, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		Err("Failed to set Main Logger: ", err)
		return err
	}
	MainLog = log.New(mLF, "MAIN:", log.Ldate|log.Ltime)
	return nil
}

func Make(prefix string, fileNames ...string) func(...interface{}) {
	f1 := makeLogger(prefix, fileNames...)
	f := func(args ...interface{}) {
		f1.Println(args...)
		MainLog.Println(args...)
	}
	return f
}

func MakeErr(prefix string, fileNames ...string) func(...interface{}) {
	f1 := makeLogger(prefix, fileNames...)
	f := func(args ...interface{}) {
		f1.Println(args...)
		ErrLog.Println(args...)
	}
	return f
}

func makeLogger(prefix string, fileNames ...string) *log.Logger {
	files := make([]io.Writer, 0)
	for _, filename := range fileNames {
		xFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			if ErrLog != nil {
				Err("Failed to make logger: ", err)
			} else {
				log.Println("Failed to make logger: ", err)
			}
			return nil
		}
		files = append(files, xFile)
	}
	Multi := io.MultiWriter(files...)
	return log.New(Multi, prefix, log.Ldate|log.Ltime)
}
