package app

import (
	"net/http"

	"github.com/cheyilin/boabot/pkg/api"
	"github.com/cheyilin/boabot/pkg/boa/aimba"
	"github.com/cheyilin/boabot/pkg/zoom"
)

// Handler is the API entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	api.NewHandler(ZoomResponser)(w, r)
}

// ZoomResponser returns BoA response in Zoom message format
// https://marketplace.zoom.us/docs/guides/chatbots/sending-messages
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, api.Error(http.StatusMethodNotAllowed)
	}

	accessToken, err := zoom.GetAccessToken()
	if err != nil {
		return nil, api.Error(http.StatusInternalServerError)
	}

	cmd, err := zoom.ParseCommand(r)
	if err != nil {
		return nil, api.Error(http.StatusInternalServerError)
	}

	resp := &zoom.Response{
		RobotJID:  zoom.Conf.RobotJID,
		ToJID:     cmd.Payload.UserJID,
		AccountID: cmd.Payload.AccountID,
		Content: &zoom.Content{
			Head: &zoom.Head{
				Text: aimba.Boa.GetAnswer(),
			},
		},
	}

	err = zoom.SendMessage(accessToken, resp)
	if err != nil {
		return nil, api.Error(http.StatusInternalServerError)
	}

	return resp, nil
}
