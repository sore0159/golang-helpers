package jsend

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func Read(r *http.Request, j interface{}) (err error) {
	var bytes []byte
	bytes, err = ioutil.ReadAll(io.LimitReader(r.Body, MAXSIZE))
	if my, bad := Check(err, "read body readall"); bad {
		return my
	}
	err = r.Body.Close()
	if my, bad := Check(err, "read body close"); bad {
		return my
	}
	err = json.Unmarshal(bytes, j)
	if my, bad := Check(err, "read unmarshal"); bad {
		return my
	}
	return nil
}
