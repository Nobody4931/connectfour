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
	num := 0

	var plr Space
	fmt.Scanf("%d\n", &plr)

	for !game.IsGameOver() {
		for _, row := range game.Board {
			fmt.Println(row)
		}

		if turn == plr /* && num > 5 */ {
			if optimalMove := game.Minimax(turn); optimalMove.Next != nil {
				fmt.Printf("\nRecommended move: %d\n", game.Minimax(turn).Next.Move)
			} else {
				fmt.Printf("\nyeah you probably lost good luck loser\n")
			}
		}
		num++

		var move int
		fmt.Printf("Player %d's move: ", turn)
		fmt.Scanf("%d\n", &move)

		game.Place(move, turn)
		turn ^= PlayerXor

		fmt.Println()
	}
}
