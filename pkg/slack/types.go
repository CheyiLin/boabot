package slack

// SlashCmdResponseType* reperesnt response enum types
const (
	SlashCmdResponseTypeDefault = "ephemeral"
	SlashCmdResponseTypeChannel = "in_channel"
)

// SlashCmdResponse represents a response of slack's slash command
type SlashCmdResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
