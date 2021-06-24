package aimba

import (
	"github.com/cheyilin/boabot/pkg/boa"
)

var (
	Boa = boa.DefaultBoa("aimba")
)

func init() {
	Boa.ExtendAnswers(answers)
}
