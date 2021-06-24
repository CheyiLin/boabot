package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa/aimba"
	"github.com/nlopes/slack"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	api.NewHandler(SlackResponser)(w, r)
}

// SlackResponser returns BoA response in Slack message format
func SlackResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, api.Error(http.StatusMethodNotAllowed)
	}

	cmd, err := slack.SlashCommandParse(r)
	if err != nil {
		return nil, api.Error(http.StatusBadRequest)
	}

	if cmd.Text == "" {
		cmd.Text = aimba.Boa.DefaultQuestion()
	}

	sb := &strings.Builder{}
	if cmd.UserID != "" {
		fmt.Fprintf(sb, "<@%s>\n", cmd.UserID)
	}
	fmt.Fprintf(sb, ">%s\n", cmd.Text)
	fmt.Fprintf(sb, "*%s*", aimba.Boa.GetAnswer())

	resp := &slack.Msg{
		// response in channel
		ResponseType: slack.ResponseTypeInChannel,
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				// use layout blocks for rich messages
				// https://api.slack.com/reference/block-kit/blocks
				slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: slack.MarkdownType,
						Text: sb.String(),
					},
				},
			},
		},
	}
	return resp, nil
}
