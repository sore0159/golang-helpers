package mybad

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	log.Println("TEST ONE")
}

func TestSecond(t *testing.T) {
	var testErr error
	my, bad := Check(testErr, "HELLO")
	log.Println("BAD:", bad)
	log.Println(my.MuleError())
	my, bad = Check(my, "NIL MULEERROR CHECK")
	log.Println("BAD:", bad)
	log.Println(my.MuleError())
	testErr = errors.New("TEST ERROR1")
	my, bad = Check(testErr, "HELLO", "time", func() interface{} { return time.Now() })
	log.Println("BAD:", bad)
	log.Println(my.MuleError())
	// ------------------- //
}

func TestThird(t *testing.T) {
	testErr := errors.New("TEST ERROR2")
	my, bad := Check(testErr, "HELLO2", "", "time", func() interface{} { return time.Now() })
	log.Println("BAD:", bad)
	log.Println(my)
	log.Println(my.Error())
	log.Println(my.LogError())
	log.Println(my.MuleError())
}

func TestFourth(t *testing.T) {
	testErr := errors.New("TEST ERROR3")
	my, bad := BuildCheckTest(testErr, "HELLO3", "time", func() interface{} { return time.Now() })
	log.Println("BAD:", bad)
	log.Println(my.LogError())
}

var BuildCheckTest = BuildCheck("package", "mybad", "testing", 12345)
