package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dictionary := readFileIntoDict("./word_dictionary.txt")
	guesser := NewGuesser()
	guesser.SetDictionary(dictionary)
	benchmark := NewBenchmark(guesser)
	fmt.Println("New benchmark initialised. Starting benchmark.")
	benchmark.RunBenchmark()
}

func readFileIntoDict(path string) map[string]bool {
	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	wordMap := map[string]bool{}
    for scanner.Scan() {
		word := scanner.Text()
		wordMap[word] = true
    }
	return wordMap
}