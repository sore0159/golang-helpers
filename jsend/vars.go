package jsend

import (
	"log"
)

const (
	MAXSIZE int64 = 1048576
)

var Log = log.Println

func SetLogger(f func(...interface{})) {
	Log = f
}
