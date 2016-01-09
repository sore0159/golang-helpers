package jsend

import (
	"net/http"
)

// Everything is awesome!
func Emmet(w http.ResponseWriter) {
	Success(w, nil)
}

// But I still haven't found what I'm looking for...
func U2(w http.ResponseWriter) {
	Fail(w, 404, map[string]string{"generic": "resource not found"})
}

// YOU SHALL NOT PASS!
func Gandalf(w http.ResponseWriter) {
	Fail(w, 401, map[string]string{"authorization": "access denied"})
}

// Your request was bad and you should feel bad.
func Zoidberg(w http.ResponseWriter) {
	Fail(w, 400, map[string]string{"generic": "bad specification for requested resource"})
}

// You are in error!  You did not discover your mistake,
// you have made two errors!  You are flawed and imperfect.
func Kirk(w http.ResponseWriter) {
	Error(w, "there has been a server error")
}
