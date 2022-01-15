package main

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed words.txt
var content []byte

func TestSolver(t *testing.T) {
	for _, soln := range bytes.Split(content, []byte("\n")) {
		soln := *(*[5]byte)(soln)
		solver, err := NewSolver(content)
		if err != nil {
			t.Errorf("error creating solver for %s:%s", soln, err)
		}
		solns := len(solver.words)
		for solns > 1 {
			guess := solver.GetGuess()
			result := getResult(soln, guess)
			solver.ApplyResult(guess, result)
			solns = len(solver.words)
		}
		if solns != 1 {
			t.Error("did not arrive at single solution")
		}
		if soln != solver.words[0] {
			t.Error("solutions are mismatched")
		}
	}
}

func getResult(soln [5]byte, guess [5]byte) [5]byte {
	var result [5]byte
	for i, b := range guess {
		switch {
		case soln[i] == b:
			result[i] = 50
		case contains(soln, b):
			result[i] = 49
		default:
			result[i] = 48
		}
	}
	return result
}

func contains(a [5]byte, b byte) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}
	return false
}
