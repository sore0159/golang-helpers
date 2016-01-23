package mylog

import (
	"io"
	"log"
	"mule/mybad"
	"os"
)

// StockErrorLogger just logs to stderr with normal
// flags and "ERROR" prefix
func StockErrorLogger() *Logger {
	return New(os.Stderr, "ERROR:", log.Ldate|log.Ltime)
}

// StockInfoLogger just logs to stdout with normal
// flags and "INFO" prefix
func StockInfoLogger() *Logger {
	return New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
}

// StockFileLogger just logs to the given file with normal
// flags and no prefix, returning an error on file failure
func StockFileLogger(fName string) (*Logger, error) {
	w, err := FilesWriter(fName)
	if my, bad := mybad.Check(err, "stockfilelogger failure"); bad {
		return nil, my
	}
	return New(w, "", log.Ldate|log.Ltime), nil
}

// StockErrorFileLogger logs to both stderr and the given file
// with normal flags and "ERROR" prefix, returning an error
// on file failure
func StockErrorFileLogger(fName string) (*Logger, error) {
	fw, err := FilesWriter(fName)
	if my, bad := mybad.Check(err, "stockerrorfilelogger failure"); bad {
		return nil, my
	}
	w := io.MultiWriter(fw, os.Stderr)
	return New(w, "ERROR:", log.Ldate|log.Ltime), nil
}

// MustErrFile is a quick call for setting up a global
// variable logger
func MustErrFile(fName string) *Logger {
	lg, err := StockErrorFileLogger(fName)
	if err != nil {
		panic(err.(*mybad.MuleError).MuleError())
	}
	return lg
}

func Must(lg *Logger, err error) *Logger {
	if my, bad := mybad.Check(err, "Must failure"); bad {
		panic(my.MuleError())
	}
	return lg
}
