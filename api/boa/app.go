package boa

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		status := http.StatusMethodNotAllowed
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%d %s", status, http.StatusText(status))
		return
	}

	resp := struct {
		Answer string `json:"answer"`
	}{
		Answer: getAnswer(),
	}

	jsonBs, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBs)
}

func getAnswer() string {
	answersCount := len(answers)
	if answersCount == 0 {
		return defaultAnswer
	}
	return answers[rand.Intn(answersCount)]
}

const (
	defaultAnswer = "順從你的心"
)

var (
	answers = []string{
		"為做最好的決定，保持冷靜",
		"別忘了生活還有樂趣",
		"轉移你的注意力",
		"援助將使你的發展取得成功",
		"不要指望它",
		"結果將取決於你的選擇",
		"把注意力集中在細節上",
		"試著改變你的日常生活",
		"你會因為你做了而感到快樂的",
		"絕不",
		"採用一種大膽的態度",
		"等待",
		"還有很多事等著你努力",
		"一年以後，已經不重要了",
		"跟隨別人的領導",
		"準備迎接意外",
		"是的",
		"清除障礙",
		"去遵循專家的建議",
		"懷疑它吧",
		"它可能是非凡的",
		"你不是真心的",
		"晚一點再處理",
		"負責",
		"實際一點",
		"隨心而行",
		"情況很快就會有改變",
		"你必須現在行動",
		"保持靈活",
		"是時候了",
		"晚一點再處理",
		"相信自己獨到的想法",
		"為什麼不做一個計畫",
		"看看會發生什麼",
		"可能會發生結果令人吃驚的事件",
		"要耐心",
		"別浪費你的時間",
		"這不是一件普通的事，慎重考慮",
		"絕對不行",
		"無論如何",
		"想想什麼才是重要的",
		"答案就在你的後院",
		"無關緊要，時間沖淡一切",
		"你將不得不妥協",
		"不要再問了",
		"它會影響到別人對你的看法",
		"它肯定會讓事情更有趣",
		"最好等一等",
		"現在是制定計劃的好時機",
		"你需要收集更多資訊",
		"必須的",
		"現在不要再提問題了，神有些累了",
		"如果你不抗拒，可以這麼幹",
		"最好等一等",
		"這不明智",
		"要謹慎",
	}
)
