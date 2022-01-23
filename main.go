package main

import "fmt"

func main() {
	guesser := NewGuesser()
	guesser.guessWord("___", []int{0,1,2}, []rune("ab"))
	// benchmark := NewBenchmark(guesser)
	fmt.Println("New benchmark initialised!")
	// benchmark.RunBenchmark()
}

type DummyGuesser struct{}

func(*DummyGuesser) NewGame() {}
func(*DummyGuesser) NextGuess() string {
	return "hello"
}
func(*DummyGuesser) Feedback(guess string, feedback []WordleResponse) {}


