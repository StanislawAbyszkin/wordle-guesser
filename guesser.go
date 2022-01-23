package main

import "fmt"

type Guesser struct {
	previousGuesses []string
	availableChars  []rune
	nextGuess       string
	dictionary      map[string]bool
}

func NewGuesser() *Guesser {
	return &Guesser{
		dictionary: map[string]bool{
			"hello": true,
			"anime": true,
			"magic": true,
			"bx": true,
		},
	}
}

func (g *Guesser) NewGame() {
	g.previousGuesses = []string{}
	g.availableChars = []rune{}
	g.nextGuess = ""
	for i := 97; i < 97+26; i++ {
		g.availableChars = append(g.availableChars, rune(i))
	}
	fmt.Println(g.availableChars)
}
func (g *Guesser) NextGuess() string {
	return "hello"
}
func (*Guesser) Feedback(guess string, feedback []WordleResponse) {}

func (g *Guesser) guessWord(guess string, indices []int, availableChars []rune) string {
	// fmt.Println(guess, indices)
	if len(indices) == 0 {
		return guess
	}

	for i, idx := range indices {
		remainingIndcies := make([]int, len(indices))
		copy(remainingIndcies, indices)
		remainingIndcies = remainingIndcies[i+1:]
		// fmt.Println("idx:", indices, "rem idxs:", remainingIndcies)
		for _, c := range availableChars {
			newGuess := replaceAtIndex(guess, c, idx)
			fmt.Println("i", i,"idx", idx, "indexes", indices, "newGuess", newGuess)
			output := g.guessWord(newGuess, remainingIndcies, availableChars)
			if ok := g.dictionary[output]; ok {
				return output
			}
		}
	}
	fmt.Println("returning")
	return ""
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
