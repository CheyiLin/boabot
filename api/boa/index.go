package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/cheyilin/boabot/pkg/boa"
	"github.com/cheyilin/boabot/pkg/slack"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		status := http.StatusMethodNotAllowed
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%d %s\n", status, http.StatusText(status))
		return
	}

	resp := slack.NewSlashResponse(boa.GetAnswer())
	jsonBs, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBs)
}
