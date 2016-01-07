package jsend

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func Read(r *http.Request, j interface{}) (ok bool) {
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadAll(io.LimitReader(r.Body, MAXSIZE)); err != nil {
		Log("JSEND R BODY READ ERROR:", err)
		return false
	}
	if err = r.Body.Close(); err != nil {
		Log("JSEND R BODY CLOSE ERROR:", err)
		return false
	}
	if err = json.Unmarshal(bytes, j); err != nil {
		Log("JSEND UNMARSHAL ERROR FOR:", j, err)
		return false
	}
	return true
}
