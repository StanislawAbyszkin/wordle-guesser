package main

import "fmt"

func main() {
	guesser := NewGuesser()
	benchmark := NewBenchmark(guesser)
	fmt.Println("New benchmark initialised. Starting benchmark.")
	benchmark.RunBenchmark()
}