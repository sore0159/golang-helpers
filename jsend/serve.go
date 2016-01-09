// Shell for encoding per the JSEND api spec
// http://labs.omniti.com/labs/jsend
package jsend

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, obj interface{}) {
	raw, err := json.Marshal(obj)
	if err != nil {
		Log("ERROR JSON MARSHALLING IN JSEND SUCCESS FUNC:", obj, err)
		Error(w, "ERROR MARSHALLING OBJECT TO JSON IN JSEND SUCCESS FUNC")
		return
	}
	ss := &successShell{"success", raw}
	ss.Serve(w)
}

func Error(w http.ResponseWriter, msg string) {
	es := &errorShell{"error", msg, 500}
	es.Serve(w)
	return
}

func Fail(w http.ResponseWriter, code int, data interface{}) {
	raw, err := json.Marshal(data)
	if err != nil {
		Log("ERROR JSON MARSHALLING IN JSEND FAIL FUNC:", data, err)
		Error(w, "ERROR MARSHALLING OBJECT TO JSON IN JSEND FAIL FUNC")
		return
	}
	fs := &failShell{"fail", raw, code}
	fs.Serve(w)
	return
}
