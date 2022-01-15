package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content := ReadWordList()
	solver, err := NewSolver(content)
	checkErr(err)
	solns := len(solver.words)
	for solns > 1 {
		fmt.Printf("%d possible solutions\n", solns)
		guess := solver.GetGuess()
		fmt.Printf("guess: %s\n", guess)
		fmt.Printf("input result: ")
		result, err := ReadResult()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		solver.ApplyResult(guess, result)
		solns = len(solver.words)
	}
	if solns != 1 {
		fmt.Println("no possible solution")
		os.Exit(1)
	}
	fmt.Printf("solution: %s\n", solver.words[0])
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadWordList() []byte {
	if len(os.Args) != 2 {
		fmt.Println("invalid args")
		os.Exit(1)
	}
	content, err := ioutil.ReadFile(os.Args[1])
	checkErr(err)
	return content
}

func ReadResult() ([5]byte, error) {
	var result [5]byte
	var input []byte
	fmt.Scanln(&input)
	if len(input) != 5 {
		return result, fmt.Errorf("input length %d != 5", len(input))
	}
	for i, b := range input {
		switch b {
		case 48, 49, 50:
			result[i] = b
		default:
			return result, fmt.Errorf("input character %c != 0|1|2", b)
		}
	}
	return result, nil
}

type Solver struct {
	letters  map[byte]int
	solution [5]map[byte]bool
	words    [][5]byte
}

func NewSolver(content []byte) (*Solver, error) {
	var solver Solver
	solver.letters = make(map[byte]int)
	solver.solution = [5]map[byte]bool{
		make(map[byte]bool),
		make(map[byte]bool),
		make(map[byte]bool),
		make(map[byte]bool),
		make(map[byte]bool),
	}
	for _, slice := range bytes.Split(content, []byte("\n")) {
		if len(slice) != 5 {
			return nil, fmt.Errorf("invalid word %s", slice)
		}
		var word [5]byte
		for i, b := range slice {
			word[i] = b
			solver.solution[i][b] = true
			solver.letters[b] += 1
		}
		solver.words = append(solver.words, word)
	}
	return &solver, nil
}

func (solver *Solver) ApplyResult(guess [5]byte, result [5]byte) {
	for i := 0; i < 5; i++ {
		switch result[i] {
		case 48:
			for _, m := range solver.solution {
				delete(m, guess[i])
			}
		case 49:
			delete(solver.solution[i], guess[i])
		case 50:
			solver.solution[i] = make(map[byte]bool)
			solver.solution[i][guess[i]] = true
		}
	}
	var words [][5]byte
	for _, word := range solver.words {
		valid := true
		for i, b := range word {
			if _, ok := solver.solution[i][b]; !ok {
				valid = false
			}
		}
		if valid {
			words = append(words, word)
		}
	}
	solver.words = words
}

func (solver *Solver) GetGuess() [5]byte {
	var guess [5]byte
	guessRank := 0
	for _, word := range solver.words {
		wordRank := 0
		wordLetters := make(map[byte]bool)
		for _, b := range word {
			wordLetters[b] = true
		}
		for letter := range wordLetters {
			wordRank += solver.letters[letter]
		}
		if wordRank > guessRank {
			guess = word
			guessRank = wordRank
		}
	}
	return guess
}
