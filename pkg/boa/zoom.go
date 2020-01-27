package boa

import (
	"net/http"
	// zoom "github.com/himalayan-institute/zoom-lib-golang"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

// ZoomResponser returns BoA response in Zoom message format
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	resp := &Response{Message: "hello world!"}

	return resp, nil
}
