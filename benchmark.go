package main

import (
	"fmt"
	"strings"
)

type Benchmark struct {
	guesser           WordleGuesser
	maxGuessesPerGame int
}

func NewBenchmark(guesser WordleGuesser) *Benchmark {
	return &Benchmark{
		guesser:           guesser,
		maxGuessesPerGame: 5,
	}
}

func (b *Benchmark) RunBenchmark() {
	testWords := []string{"anime", "hello"}
	testGuesses := make([]int, len(testWords))
	totalGuesses := 0
	for idx, testWord := range testWords {
		b.guesser.NewGame()
		currentTestGuesses := 0
		for i := 0; i < b.maxGuessesPerGame; i++ {
			guess := b.guesser.NextGuess()
			totalGuesses++
			currentTestGuesses++
			if testWord == guess {
				break
			}
			feedback := getFeedback(guess, testWord)
			b.guesser.Feedback(guess, feedback)
		}
		testGuesses[idx] = currentTestGuesses
	}

	fmt.Println("total guesses:", totalGuesses)
	fmt.Println("per test guesses:", testGuesses)
}

func getFeedback(guess, actual string) []WordleResponse {
	if len(guess) != len(actual) {
		panic(fmt.Sprintf("len of provided guess word is different from expected; expected %d, got %d", len(guess), len(actual)))
	}
	response := make([]WordleResponse, len(guess))

	for i := range guess {
		currentChar := guess[i]
		if currentChar == actual[i] {
			response[i] = Correct
			continue
		}
		if strings.Contains(actual, string(currentChar)) {
			response[i] = Present
			continue
		}
		response[i] = NotPresent
	}
	// fmt.Println("Response", response)
	return response
}
