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

	game.Place(1, PlayerTwo)
	game.Place(1, PlayerTwo)
	game.Place(1, PlayerTwo)
	game.Place(1, PlayerOne)

	game.Place(2, PlayerOne)
	game.Place(2, PlayerOne)

	game.Place(3, PlayerTwo)
	game.Place(3, PlayerOne)

	fmt.Println(game.Minimax(PlayerOne).Next.Move)

	/* rand.Seed(time.Now().Unix())

	turn := PlayerOne
	for i := 0; i < 6; i++ {
		move := rand.Intn(opts.Cols)
		for !game.Place(move, turn) {
			move = rand.Intn(opts.Cols)
		}
		turn ^= PlayerXor
	}

	for _, row := range game.Board {
		fmt.Println(row)
	}

	for !game.IsGameOver() {
		prediction := game.Minimax(turn).Next.Move
		game.Place(prediction, turn)
		fmt.Printf("%d: %d\n", turn, prediction)
		turn ^= PlayerXor
	} */
}
