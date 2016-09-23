package users

import (
	"fmt"
	"log"
	"testing"
)

func TestFirst(t *testing.T) {
	log.Println("TEST ONE")
}

func TestSecond(t *testing.T) {
	rg, ok := GetRegistry()
	if !ok {
		log.Println("FAILED GETREGISTRY")
		return
	}
	log.Println("GOT REGISTRY")

	if rg.validLogin("testone", "testpw") {
		log.Println("LOGIN 1 VALID")
	} else {
		log.Println("LOGIN 1 INVALID")
	}

	nameOk, pwOk, dbOk := rg.createUser("testone", "testpw")
	fmt.Println("CREATE BOOLS:", nameOk, pwOk, dbOk)

	if rg.validLogin("testone", "testpw") {
		log.Println("LOGIN 2 VALID")
	} else {
		log.Println("LOGIN 2 INVALID")
	}
}

// CREATE TABLE userinfo (
//		name varchar(20) PRIMARY KEY,
//		password varchar(20)
// );
