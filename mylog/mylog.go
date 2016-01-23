package mylog

import (
	"log"
)

var (
	DEVNULL = &Nuller{}
	NULLLOG = &Logger{log.New(DEVNULL, "", 0), DEVNULL}
)

type Nuller struct {
}

func (*Nuller) Write(p []byte) (n int, err error) {
	return len(p), nil
}
