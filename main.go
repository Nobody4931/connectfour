package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func main() {
	// TEST CODE, REMOVE LATER
	opts := NewOptions()
	game := NewGame(&opts)
	turn := PlayerOne

	for !game.IsGameOver() {
		for _, row := range game.Board {
			fmt.Println(row)
		}

		if optimalMove := game.Minimax(turn); optimalMove.Next != nil {
			fmt.Printf("\nRecommended move: %d\n", game.Minimax(turn).Next.Move)
		} else {
			fmt.Printf("\nyeah you probably lost good luck loser\n")
		}

		var move int
		fmt.Printf("Player %d's move: ", turn)
		fmt.Scanf("%d\n", &move)

		game.Place(move, turn)
		turn ^= PlayerXor

		fmt.Println()
	}
}
