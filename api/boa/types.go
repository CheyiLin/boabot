package boa

// SlackSlashCmdResponseType* reperesnt response enum types
const (
	SlackSlashCmdResponseTypeDefault = "ephemeral"
	SlackSlashCmdResponseTypeChannel = "in_channel"
)

// SlackSlashCmdResponse represents a response of slack's slash command
type SlackSlashCmdResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
