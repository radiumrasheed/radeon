package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string {
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
	}

	board[0][0] = "X"
	board[2][0] = "0"
	board[1][1] = "X"
	board[2][1] = "0"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], ""))
	}
}
