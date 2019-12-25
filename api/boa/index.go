package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/cheyilin/boabot/pkg/boa"
	"github.com/nlopes/slack"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		respStatus(w, http.StatusMethodNotAllowed)
		return
	}

	cmd, err := slack.SlashCommandParse(r)
	if err != nil {
		respStatus(w, http.StatusBadRequest)
		return
	}

	sb := &strings.Builder{}
	if cmd.UserID != "" {
		fmt.Fprintf(sb, "<@%s>! ", cmd.UserID)
	}
	if cmd.Text != "" {
		fmt.Fprintf(sb, "Q: %s", cmd.Text)
	}
	fmt.Fprintf(sb, "A: %s", boa.GetAnswer())

	resp := &slack.Msg{
		ResponseType: slack.ResponseTypeInChannel,
		Text:         sb.String(),
	}
	jsonBs, err := json.Marshal(resp)
	if err != nil {
		respStatus(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBs)
}

func respStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%d %s\n", status, http.StatusText(status))
}
