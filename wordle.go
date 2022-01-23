package main

type WordleResponse int

const (
	NotPresent WordleResponse = iota
	Present
	Correct
)

type WordleGuesser interface {
	// NewGame indicates that a new word is to be guessed. This method usually would clear the state and prepare guesser to try and guess a new word.
	NewGame()
	// NextGuess expects a new guess provided by the guesser.
	NextGuess() string
	// Feedback function is called to provide guesser with feedback on how it performed guessing a word. 
	Feedback(guess string, feedback []WordleResponse)
}
