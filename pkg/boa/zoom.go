package boa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ZoomCommand struct {
	Event   string  `json:"event,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

type Payload struct {
	AccountID   string `json:"accountId,omitempty"`
	ChannelName string `json:"channelName,omitempty"`
	Cmd         string `json:"cmd,omitempty"`
	Name        string `json:"name,omitempty"`
	RobotJID    string `json:"robotJid,omitempty"`
	Timestamp   int    `json:"timestamp,omitempty"`
	ToJID       string `json:"toJid,omitempty"`
	UserID      string `json:"userId,omitempty"`
	UserJID     string `json:"userJid,omitempty"`
	UserName    string `json:"userName,omitempty"`
}

// Response - chatbot reponse
type Response struct {
	RobotJID  string   `json:"robot_jid,omitempty"`
	ToJID     string   `json:"to_jid,omitempty"`
	AccountID string   `json:"account_id,omitempty"`
	Content   *Content `json:"content,omitempty"`
}

type Content struct {
	Head *Head `json:"head,omitempty"`
}

type Head struct {
	Text string `json:"text,omitempty"`
}

// ZoomResponser returns BoA response in Zoom message format
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	cmd, err := ZoomCommandParse(r)
	if err != nil {
		return nil, Error(http.StatusBadRequest)
	}

	if cmd.Payload.Cmd == "" {
		cmd.Payload.Cmd = defaultQuestion
	}

	resp := &Response{
		RobotJID:  "v1cig4wtzjqwee4xagwkplaa@xmpp.zoom.us",
		ToJID:     cmd.Payload.UserJID,
		AccountID: cmd.Payload.AccountID,
		Content: &Content{
			Head: &Head{
				Text: GetAnswer(),
			},
		},
	}

	fmt.Println("cmd.Payload.AccountID:" + cmd.Payload.AccountID)
	fmt.Println("cmd.Payload.ChannelName:" + cmd.Payload.ChannelName)
	fmt.Println("cmd.Payload.Cmd:" + cmd.Payload.Cmd)
	fmt.Println("cmd.Payload.Name:" + cmd.Payload.Name)
	fmt.Println("cmd.Payload.RobotJID:" + cmd.Payload.RobotJID)
	fmt.Println("cmd.Payload.ToJID:" + cmd.Payload.ToJID)
	fmt.Println("cmd.Payload.UserID:" + cmd.Payload.UserID)
	fmt.Println("cmd.Payload.UserJID:" + cmd.Payload.UserJID)
	fmt.Println("cmd.Payload.UserName:" + cmd.Payload.UserName)

	fmt.Println("resp.RobotJID" + resp.RobotJID)
	fmt.Println("resp.ToJID" + resp.ToJID)
	fmt.Println("resp.AccountID" + resp.AccountID)
	fmt.Println("resp.Content.Head.Text" + resp.Content.Head.Text)

	return resp, nil
}

// ZoomCommandParse will parse the request of the zoom command
func ZoomCommandParse(r *http.Request) (z ZoomCommand, err error) {
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&z); err != nil {
		panic(err)
	}

	return z, nil
}
