package zoom

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa"
)

// BoaResponser returns BoA response in Zoom message format
// https://marketplace.zoom.us/docs/guides/chatbots/sending-messages
func BoaResponser(b boa.Boaer) api.ResponseFunc {
	return func(r *http.Request) (interface{}, error) {
		switch r.Method {
		case http.MethodGet, http.MethodPost:
			// allowed methods
		default:
			return nil, api.Error(http.StatusMethodNotAllowed)
		}

		accessToken, err := getAccessToken()
		if err != nil {
			return nil, api.Error(http.StatusInternalServerError)
		}

		cmd, err := parseCommand(r)
		if err != nil {
			return nil, api.Error(http.StatusInternalServerError)
		}

		resp := &Response{
			RobotJID:  conf.RobotJID,
			ToJID:     cmd.Payload.UserJID,
			AccountID: cmd.Payload.AccountID,
			Content: &Content{
				Head: &Head{
					Text: b.GetAnswer(),
				},
			},
		}

		err = sendMessage(accessToken, resp)
		if err != nil {
			return nil, api.Error(http.StatusInternalServerError)
		}

		return resp, nil
	}
}
