package app

import (
	"github.com/cheyilin/boabot/pkg/boa"
)

var (
	// Handler is the API entrypoint
	Handler = boa.NewHandler(boa.SlackResponser)
)
