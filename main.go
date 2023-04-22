package main

import "fmt"

func main() {
	// TEST CODE, REMOVE LATER
	opts := NewOptions()
	game := NewGame(&opts)

	game.Place(0, PlayerTwo)
	game.Place(0, PlayerTwo)
	game.Place(0, PlayerTwo)
	game.Place(0, PlayerOne)

	game.Place(1, PlayerTwo)
	game.Place(1, PlayerTwo)
	game.Place(1, PlayerOne)

	game.Place(2, PlayerTwo)
	game.Place(2, PlayerOne)

	game.Place(3, PlayerOne)

	fmt.Println(game.IsGameOver())
	fmt.Println(game.Winner() == PlayerOne)
}
