package main

import (
	"encoding/json"
	"net/http"
)

// reportError writes back errors to an HTTP client as a JSON response
// and sets the HTTP response code
func reportError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	e := ErrorMsg{Msg: err.Error()}
	out, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		w.Write(out)
	}
}
