// Shell for encoding per the JSEND api spec
// http://labs.omniti.com/labs/jsend
package jsend

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, obj interface{}) (err error) {
	raw, err := json.Marshal(obj)
	if my, bad := Check(err, "success marshal", "object", obj); bad {
		Error(w, my.Error())
		return my
	}
	ss := &successShell{"success", raw}
	return ss.Serve(w)
}

func Error(w http.ResponseWriter, msg string) (err error) {
	es := &errorShell{"error", msg, 500}
	return es.Serve(w)
}

func Fail(w http.ResponseWriter, code int, data interface{}) (err error) {
	raw, err := json.Marshal(data)
	if my, bad := Check(err, "fail marshal", "data", data); bad {
		Error(w, my.Error())
		return my
	}
	fs := &failShell{"fail", raw, code}
	return fs.Serve(w)
}
