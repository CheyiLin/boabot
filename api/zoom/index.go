package app

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/boa"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	boa.NewHandler(boa.ZoomResponser)(w, r)
}
