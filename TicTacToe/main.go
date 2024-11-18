package main

import (
	"fmt"
	"os"
)

var board [3][3]string
var currentPlayer string

// Initialize the game board
func initBoard() {
	board = [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	currentPlayer = "X" // Start with player "X"
}

// Print the game board
func printBoard() {
	fmt.Println("Current board:")
	for _, row := range board {
		fmt.Println(row)
	}
}

// Check if a player has won
func checkWin(player string) bool {
	// Check rows, columns, and diagonals
	for i := 0; i < 3; i++ {
		// Check rows
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		// Check columns
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

// Check if the game is a draw
func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false // There's still an empty space, so it's not a draw
			}
		}
	}
	return true
}

// Switch the current player
func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

// Handle player input
func handleInput() {
	var row, col int
	for {
		fmt.Printf("Player %s, enter row and column (0-2): ", currentPlayer)
		_, err := fmt.Scanf("%d %d", &row, &col)

		// If the input is invalid or out of bounds, prompt again
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != " " {
			fmt.Println("Invalid input, please try again.")
			// Clear the buffer after incorrect input
			fmt.Scanln()
		} else {
			// Place the mark on the board
			board[row][col] = currentPlayer
			break
		}
	}
}

func main() {
	// Initialize the game
	initBoard()

	// Game loop
	for {
		printBoard()

		// Handle player input
		handleInput()

		// Check if the current player has won
		if checkWin(currentPlayer) {
			printBoard()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		// Check if the game is a draw
		if checkDraw() {
			printBoard()
			fmt.Println("It's a draw!")
			break
		}

		// Switch to the other player
		switchPlayer()
	}

	// Prompt for exiting
	fmt.Println("Game over! Press Enter to exit.")
	fmt.Scanln()
	os.Exit(0)
}
