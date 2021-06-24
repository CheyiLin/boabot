package kk

import (
	"github.com/cheyilin/boabot/pkg/boa"
)

var (
	Boa = boa.DefaultBoa("kk")
)

func init() {
	Boa.ExtendAnswers(answers)
}
