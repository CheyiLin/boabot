package boa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

	accessToken, err := getAccessToken()
	if err != nil {
		return nil, Error(http.StatusInternalServerError)
	}

	cmd, err := ZoomCommandParse(r)
	if err != nil {
		return nil, Error(http.StatusInternalServerError)
	}

	resp := &Response{
		RobotJID:  conf.RobotJID,
		ToJID:     cmd.Payload.UserJID,
		AccountID: cmd.Payload.AccountID,
		Content: &Content{
			Head: &Head{
				Text: GetAnswer(),
			},
		},
	}

	// debug
	fmt.Println(resp)
	//

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
		fmt.Printf("[Error] Parse zoom commad decoder: %v", err)
		return z, err
	}

	return z, nil
}

func getAccessToken() (string, error) {
	url := "https://api.zoom.us/oauth/token?grant_type=client_credentials"

	b := base64.StdEncoding.EncodeToString([]byte(conf.ClientID + ":" + conf.ClientSecret))

	m := make(map[string]string)
	m["authorization"] = "Basic " + b
	m["Content-Type"] = "application/json"

	resp, err := httpPostRequest(url, nil, m)
	if err != nil {
		fmt.Printf("[Error] Get access token: %v", err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var accessTokenResponse AccessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)

	// debug
	fmt.Println(accessTokenResponse.AccessToken)
	//

	return accessTokenResponse.AccessToken, nil
}

func sendMessage(accessToken string, r *Response) error {
	url := "https://api.zoom.us/v2/im/chat/messages"

	j, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("[Error] Send message json marshal: %v", err)
		return err
	}

	m := make(map[string]string)
	m["authorization"] = "Bearer " + accessToken
	m["Content-Type"] = "application/json"

	_, err = httpPostRequest(url, bytes.NewBuffer(j), m)
	if err != nil {
		fmt.Printf("[Error] Send message: %v", err)
		return err
	}

	return nil
}

func httpPostRequest(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
