package jsend

import (
	"log"
	"net/http"
	"testing"
)

func TestFirst(t *testing.T) {
	log.Println("TEST FIRST")
}

func XTestSecond(t *testing.T) {
	const SERVPORT = ":8080"
	log.Println("STARTING SERVER AT", SERVPORT)
	http.HandleFunc("/serve", pageTestServeJson)
	http.HandleFunc("/read", pageTestReadJson)
	err := http.ListenAndServe(SERVPORT, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("STOPPING SERVER")
}
func pageTestServeJson(w http.ResponseWriter, r *http.Request) {
	g := &struct{}{}
	Success(w, g)
}

func pageTestReadJson(w http.ResponseWriter, r *http.Request) {
	g := &struct{}{}
	if my, bad := Check(Read(r, g), "test read page"); bad {
		log.Println("ERROR READING", g, ":", my.MuleError())
	}
	log.Println("READ", g)
}
