package app

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/boa"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	boa.NewHandler(ZoomResponser)(w, r)
}

// ZoomResponser returns BoA response in Zoom message format
// https://marketplace.zoom.us/docs/guides/chatbots/sending-messages
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, boa.Error(http.StatusMethodNotAllowed)
	}

	accessToken, err := getAccessToken()
	if err != nil {
		return nil, boa.Error(http.StatusInternalServerError)
	}

	cmd, err := parseCommand(r)
	if err != nil {
		return nil, boa.Error(http.StatusInternalServerError)
	}

	resp := &Response{
		RobotJID:  conf.RobotJID,
		ToJID:     cmd.Payload.UserJID,
		AccountID: cmd.Payload.AccountID,
		Content: &Content{
			Head: &Head{
				Text: boa.GetAnswer(),
			},
		},
	}

	err = sendMessage(accessToken, resp)
	if err != nil {
		return nil, boa.Error(http.StatusInternalServerError)
	}

	return resp, nil
}
