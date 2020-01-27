package boa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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

// AccessTokenResponse - chatbot reponse
type AccessTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	Scope       string `json:"scope,omitempty"`
}

// ZoomResponser returns BoA response in Zoom message format
// https://marketplace.zoom.us/docs/guides/chatbots/sending-messages
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	accessToken := getAccessToken()

	cmd, err := ZoomCommandParse(r)
	if err != nil {
		return nil, Error(http.StatusBadRequest)
	}

	resp := &Response{
		RobotJID:  os.Getenv("ROBOT_JID"),
		ToJID:     cmd.Payload.UserJID,
		AccountID: cmd.Payload.AccountID,
		Content: &Content{
			Head: &Head{
				Text: GetAnswer(),
			},
		},
	}

	err = sendMessage(accessToken, resp)
	if err != nil {
		return nil, Error(http.StatusInternalServerError)
	}

	return resp, nil
}

// ZoomCommandParse will parse the request of the zoom command
func ZoomCommandParse(r *http.Request) (z ZoomCommand, err error) {
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&z); err != nil {
		return z, err
	}

	return z, nil
}

func getAccessToken() string {
	url := "https://api.zoom.us/oauth/token?grant_type=client_credentials"

	b := base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID") + ":" + os.Getenv("CLIENT_SECRET")))
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("authorization", "Basic "+b)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var accessTokenResponse AccessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)

	return accessTokenResponse.AccessToken
}

func sendMessage(accessToken string, r *Response) error {
	url := "https://api.zoom.us/v2/im/chat/messages"

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}
