package slack

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa"
	"github.com/nlopes/slack"
)

// BoaResponser returns BoA response in Slack message format
func BoaResponser(b boa.Boaer) api.ResponseFunc {
	return func(r *http.Request) (interface{}, error) {
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
			cmd.Text = b.DefaultQuestion()
		}

		sb := &strings.Builder{}
		if cmd.UserID != "" {
			fmt.Fprintf(sb, "<@%s>\n", cmd.UserID)
		}
		fmt.Fprintf(sb, ">%s\n", cmd.Text)
		fmt.Fprintf(sb, "*%s*", b.GetAnswer())

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
}
