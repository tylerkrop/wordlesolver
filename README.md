# wordlesolver

## Background

I wrote this program to solve puzzles from https://www.powerlanguage.co.uk/wordle/

## Running

The program can be run directly using `go` via `go run . <path-to-word-list>`.
Alternatively, the binary `wordlesolver` can be built using `go build`.

## User Guide

The program requires a single positional argument which is the path to a newline separated word file. The word file should only contain 5 letter words.

The program will prompt the to input a guess and type the result of the guess. The result should be input in the same order as the letters in the guess. The valid result input is as follows:

|character|meaning|
|-|-|
|0|letter is absent in the solution|
|1|letter is present in the solution, but not in the correct index|
|2|letter is present and in the correct index|

The program terminates when either a correct solution is found or no possible solutions remain.
