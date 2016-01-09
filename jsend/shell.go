package jsend

import (
	"encoding/json"
	"net/http"
)

type successShell struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func (ss *successShell) Serve(w http.ResponseWriter) {
	ss.Status = "success"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ss); err != nil {
		Log("JSON SERVESHELL ERROR ENCODE ERROR:", ss, "\n", err)
	}
}

type failShell struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	code   int
}

func (fs *failShell) Serve(w http.ResponseWriter) {
	fs.Status = "fail"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch fs.code {
	case 401, 404:
		w.WriteHeader(fs.code)
	default:
		w.WriteHeader(400)
	}
	if err := json.NewEncoder(w).Encode(fs); err != nil {
		Log("JSON SERVESHELL ERROR ENCODE ERROR:", fs, "\n", err)
	}
}

type errorShell struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

func (es *errorShell) Serve(w http.ResponseWriter) {
	es.Status = "error"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(es); err != nil {
		Log("JSON SERVESHELL ERROR ENCODE ERROR:", es, "\n", err)
	}
}
