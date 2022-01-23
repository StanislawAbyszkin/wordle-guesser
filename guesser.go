package main

import "fmt"

type Guesser struct {
	previousGuesses []string
	availableChars []rune
	nextGuess string
	dictionary map[string]bool
}

func NewGuesser() WordleGuesser {
	return &Guesser{
		dictionary: map[string]bool{
			"hello": true, 
			"anime": true, 
			"magic": true,
		},
	}
}

func(g *Guesser) NewGame() {
	g.previousGuesses = []string{}
	g.availableChars = []rune{}
	g.nextGuess = ""
	for i := 97; i < 97+26; i++ {
		g.availableChars = append(g.availableChars, rune(i))
	}
	fmt.Println(g.availableChars)
}
func(g *Guesser) NextGuess() string {
	return "hello"
}
func(*Guesser) Feedback(guess string, feedback []WordleResponse) {}

func (g *Guesser) guessWord(guess string, indices []int, availableChars []rune) string{
	if len(indices) == 0 {
		fmt.Println(guess)
		return guess
	}

	for i, idx := range indices {
		remainingIndcies := indices[i:]
		for _, c := range availableChars {
			newGuess := replaceAtIndex(guess, c, idx)
			output := g.guessWord(newGuess, remainingIndcies, availableChars)
			if ok := g.dictionary[output]; ok {
				return output
			}
		}
	}
	panic("word not found")
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}