package main

import "fmt"

type Guesser struct {
	previousGuesses []string
	availableChars  map[rune]bool
	nextGuess       string
	nextIndices     []int
	dictionary      map[string]bool
	totalRecCalls   int
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
	g.availableChars = map[rune]bool{}
	g.nextGuess = "-----"
	g.nextIndices = []int{0, 1, 2, 3, 4}
	for i := 97; i < 97+26; i++ {
		g.availableChars[rune(i)] = true
	}
}
func (g *Guesser) NextGuess() string {
	defer func() { fmt.Println("total calls in try:", g.totalRecCalls) }()
	return g.guessWord(g.nextGuess, g.nextIndices)
}
func (g *Guesser) Feedback(guess string, feedback []WordleResponse) {
	for idx, outcome := range feedback {
		correct := false
		char := rune(guess[idx])
		switch outcome {
		case NotPresent:
			if _, ok := g.availableChars[char]; ok {
				delete(g.availableChars, char)
			}
		case Correct:
			g.nextGuess = replaceAtIndex(g.nextGuess, char, idx)
			correct = true
		}
		if !correct {
			g.nextIndices = append(g.nextIndices, idx)
		}
	}
}

func (g *Guesser) guessWord(guess string, indices []int) string {
	g.totalRecCalls++
	if len(indices) == 0 {
		return guess
	}

	remainingIndcies := make([]int, len(indices))
	copy(remainingIndcies, indices)
	remainingIndcies = remainingIndcies[1:]
	for c := range g.availableChars {
		newGuess := replaceAtIndex(guess, c, indices[0])
		output := g.guessWord(newGuess, remainingIndcies)

		if _, ok := g.dictionary[output]; ok {
			return output
		}
	}
	return "-----"
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
