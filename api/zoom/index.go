package app

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa/aimba"
	"github.com/cheyilin/boabot/pkg/zoom"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	api.NewHandler(zoom.BoaResponser(aimba.Boa))(w, r)
}
