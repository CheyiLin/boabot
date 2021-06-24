package boa

import (
	"os"
)

type Conf struct {
	RobotJID     string
	ClientID     string
	ClientSecret string
}

var conf Conf

func init() {
	conf.RobotJID = os.Getenv("ROBOT_JID")
	conf.ClientID = os.Getenv("CLIENT_ID")
	conf.ClientSecret = os.Getenv("CLIENT_SECRET")
}
