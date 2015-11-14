package mylog

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	log.Println("TESTING")
}

func TestLogger(t *testing.T) {
	l := Make("MAKETEST", "test2.txt")
	x := SetMain("test.txt")
	y := SetErr("ertest.txt")
	log.Println("TESTING ERRORS:", x, y)
	Log("TEST")
	Err("TEST2")
	l("TEST3")
}
