package mylog

import (
	"errors"
	"log"
	"mule/mybad"
	"testing"
)

func Test1(t *testing.T) {
	log.Println("TEST ONE")
}

func Test2(t *testing.T) {
	lg := NewLogger()
	lg.Println("Shouldn't see this")
	lg2, err := lg.AddFiles("TESTFILE.txt")
	if my, bad := mybad.Check(err, "test2 failure"); bad {
		log.Println(my.MuleError())
		return
	}
	lg2.Println("Should see this")
	lg2.Ping()
	lg3 := lg2.AddStdout()
	lg3.Ping()
	testErr := errors.New("TESTERR")
	if my, bad := mybad.Check(testErr, "test2 check"); bad {
		lg3.Println(my)
	}
}
