package main

type MoveNode struct {
	Move int
	Next *MoveNode
	Prev *MoveNode
}

func (game *Game) Minimax(player Space) []int {
	panic("unimplemented")
}

func (game *Game) minimax(player Space, node MoveNode, depth int, bestMax int, bestMin int, isMaximizing bool) int {
	panic("unimplemented")
}
