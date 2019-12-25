package boa

import (
	"fmt"
	"net/http"
)

// Error is a error type wraps HTTP status code
type Error int

func (s Error) Error() string {
	status := int(s)
	return fmt.Sprintf("%d %s\n", status, http.StatusText(status))
}
