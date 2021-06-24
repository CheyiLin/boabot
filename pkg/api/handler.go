package api

import (
	"encoding/json"
	"net/http"
)

// ResponseFunc is an adapter to allow the use of ordinary functions as BoA response handlers
type ResponseFunc func(r *http.Request) (interface{}, error)

// NewHandler returns a handler func that processes response from f
func NewHandler(f ResponseFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := f(r)
		if err != nil {
			respError(w, err)
			return
		}

		jsonBs, err := json.Marshal(resp)
		if err != nil {
			respError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBs)
	}
}

func respError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	if berr, ok := err.(Error); ok {
		http.Error(w, berr.Error(), int(berr))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
