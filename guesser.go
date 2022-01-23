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
	return g.guessWord("_____", []int{0,1,2,3,4}, g.availableChars)
}
func (*Guesser) Feedback(guess string, feedback []WordleResponse) {}

func (g *Guesser) guessWord(guess string, indices []int, availableChars []rune) string {
	if len(indices) == 0 {
		// fmt.Println(guess, indices)
		return guess
	}

	remainingIndcies := make([]int, len(indices))
	copy(remainingIndcies, indices)
	remainingIndcies = remainingIndcies[1:]
	for _, c := range availableChars {
		newGuess := replaceAtIndex(guess, c, indices[0])
		output := g.guessWord(newGuess, remainingIndcies, availableChars)
		
		if ok := g.dictionary[output]; ok {
			return output
		}
	}
	return "xxxxx"
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
