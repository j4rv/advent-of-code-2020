package main

import (
	"log"
)

var exampleData = []int{0, 3, 6}
var data = []int{9, 19, 1, 6, 0, 5, 4} // just copy-paste the input into the brackets

func main() {
	input := data
	log.Println("Part One solution:", memoryGame(input, 2020))
	log.Println("Part Two solution:", memoryGame(input, 30000000))
}

func memoryGame(input []int, turns int) (curr int) {
	memory := make(map[int]int)
	// starting numbers
	for i, startingNum := range input {
		curr = startingNum
		memory[curr] = i + 1 // first turn is 1 not 0
	}
	// memory game
	for prevTurn := len(input); prevTurn < turns; prevTurn++ {
		currSpokenAtTurn, ok := memory[curr]
		memory[curr] = prevTurn
		if ok {
			curr = prevTurn - currSpokenAtTurn
		} else {
			curr = 0
		}
	}
	return curr // the number at the last turn
}
