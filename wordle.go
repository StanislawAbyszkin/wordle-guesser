package main

type WordleResponse int

const (
	NotPresent WordleResponse = iota
	Present
	Correct
)

type WordleGuesser interface {
	NewGame()
	NextGuess() string
	Feedback(guess string, feedback []WordleResponse)
}
