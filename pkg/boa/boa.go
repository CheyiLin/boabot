package boa

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	defaultQuestion = "(謎之音)"
	defaultAnswer   = "順從你的心"
)

type DefaultBoa string

func (s DefaultBoa) GetAnswer() string {
	answersCount := len(answers)
	if answersCount == 0 {
		return defaultAnswer
	}
	return answers[rand.Intn(answersCount)]
}

func (s DefaultBoa) Answers() []string {
	return answers
}

func (s DefaultBoa) ExtendAnswers(ext []string) {
	answers = append(answers, ext...)
}

func (s DefaultBoa) DefaultQuestion() string {
	return defaultQuestion
}
