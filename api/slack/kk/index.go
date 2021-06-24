package app

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa/kk"
	"github.com/cheyilin/boabot/pkg/slack"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	api.NewHandler(slack.BoaResponser(kk.Boa))(w, r)
}
