package slack

import (
	slackApi "github.com/nlopes/slack"
)

// NewSlashResponse returns a new slack response message
func NewSlashResponse(text string) *slackApi.Msg {
	return &slackApi.Msg{
		ResponseType: slackApi.ResponseTypeInChannel,
		Text:         text,
	}
}
