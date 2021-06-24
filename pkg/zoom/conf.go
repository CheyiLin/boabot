package zoom

import (
	"os"
)

var conf struct {
	RobotJID     string
	ClientID     string
	ClientSecret string
}

func init() {
	conf.RobotJID = os.Getenv("ROBOT_JID")
	conf.ClientID = os.Getenv("CLIENT_ID")
	conf.ClientSecret = os.Getenv("CLIENT_SECRET")
}
