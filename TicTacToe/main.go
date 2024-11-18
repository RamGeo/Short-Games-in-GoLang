package main

import "fmt"

var board [3][3]string

func printBoard() {
	fmt.Println("Current board:")
	for _, row := range board {
		fmt.Println(row)
	}
}

func main() {
	board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	printBoard()
	// Implement the rest of the game here.
}
