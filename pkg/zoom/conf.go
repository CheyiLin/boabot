package zoom

import (
	"os"
)

var Conf struct {
	RobotJID     string
	ClientID     string
	ClientSecret string
}

func init() {
	Conf.RobotJID = os.Getenv("ROBOT_JID")
	Conf.ClientID = os.Getenv("CLIENT_ID")
	Conf.ClientSecret = os.Getenv("CLIENT_SECRET")
}
