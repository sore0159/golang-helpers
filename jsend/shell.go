package jsend

import (
	"encoding/json"
	"net/http"
)

type successShell struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func (ss *successShell) Serve(w http.ResponseWriter) (err error) {
	ss.Status = "success"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ss)
}

type failShell struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	code   int
}

func (fs *failShell) Serve(w http.ResponseWriter) (err error) {
	fs.Status = "fail"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch fs.code {
	case 401, 404:
		w.WriteHeader(fs.code)
	default:
		w.WriteHeader(400)
	}
	return json.NewEncoder(w).Encode(fs)
}

type errorShell struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

func (es *errorShell) Serve(w http.ResponseWriter) (err error) {
	es.Status = "error"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(es)
}
