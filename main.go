package main

import "fmt"

func main() {
	guesser := NewGuesser()
	fmt.Println("New benchmark initialised. Starting benchmark.")
	benchmark := NewBenchmark(guesser)
	benchmark.RunBenchmark()
}

type DummyGuesser struct{}

func (*DummyGuesser) NewGame() {}
func (*DummyGuesser) NextGuess() string {
	return "hello"
}
func (*DummyGuesser) Feedback(guess string, feedback []WordleResponse) {}
