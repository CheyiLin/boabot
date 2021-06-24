package boa

type Boaer interface {
	GetAnswer() string
	Answers() []string
	ExtendAnswers([]string)
	DefaultQuestion() string
}
