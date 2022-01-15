# wordlesolver

## Background

I wrote this program to solve puzzles from https://www.powerlanguage.co.uk/wordle/

## Running

The program can be run directly via `go run . <path-to-word-list>`.
Alternatively, the binary `wordlesolver` can be built using `go build`.

The test can be run via `go test`. It ensures that the program will eventually find a solution for each word in the word list. The test takes around 155s to run on my machine.

## User Guide

The program requires a single positional argument which is the path to a newline separated word file. The word file should only contain 5 letter words.

The program will prompt the to input a guess and type the result of the guess. The result should be input in the same order as the letters in the guess. The valid result input is as follows:

|character|meaning|
|-|-|
|0|letter is absent in the solution|
|1|letter is present in the solution, but not in the correct index|
|2|letter is present and in the correct index|

The program terminates when either a correct solution is found or no possible solutions remain.

### Example

The following is the output from running the program against the solution "panic".

```
12972 possible solutions
guess: aeros
input result: 10000
1378 possible solutions
guess: dital
input result: 01010
332 possible solutions
guess: uncia
input result: 01121
23 possible solutions
guess: panic
input result: 22222
solution: panic
```

## Assumptions

- The program assumes a word list that only contains 5 letters. This could be changed, but I wanted to keep it simple.
- The program does not keep track of the number of guesses. From my testing, it will guess the correct solution in 6 or less guesses 87% of the time.
