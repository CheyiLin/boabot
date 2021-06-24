package aimba

import (
	"github.com/cheyilin/boabot/pkg/boa"
)

var (
	Boa     = boa.DefaultBoa("aimba")
	answers = []string{
		"你冷靜",
		"在有跟沒有之間",
		"尋求援助吧",
		"不要指望 TA",
		"魔鬼藏在細節裡",
		"BJ4",
		"就做吧",
		"毋湯",
		"大膽一點吧",
		"你就等",
		"放眼世界，征服宇宙",
		"可以問你媽",
		"有 87% 可能性",
		"謀恩丟",
		"如果你能一個禮拜不喝飲料的話",
		"可以去問石頭",
		"結果可能還不錯",
		"了不起，負責",
		"你想幹嘛就幹嘛",
		"拖延一下沒關係",
		"可能會發生令人吃驚的事",
		"答案就在你身後",
		"這種小事可以不用問",
		"如果你不討厭，就做吧",
	}
)

func init() {
	Boa.ExtendAnswers(answers)
}
