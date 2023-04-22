package main

import "math"

type MoveNode struct {
	Move int
	Next *MoveNode
	Prev *MoveNode
}

func (game *Game) Minimax(player Space) MoveNode {
	// Calculate an optimal depth to search based on how many available moves exist
	availableCols := 0
	for col := 0; col < game.Opts.Cols; col++ {
		if game.CanPlace(col) {
			availableCols++
		}
	}
	maxDepth := intLog(availableCols, 50000000)

	rootMove := MoveNode{ Move: -1 }
	game.minimax(player, &rootMove, maxDepth, math.MinInt, math.MaxInt, true)
	return rootMove
}


func (game *Game) minimax(player Space, node *MoveNode, depth int, bestMax int, bestMin int, isMaximizing bool) int {
	// TODO: Optimize this? Calls game.getAllConsecutives() twice
	if depth == 0 || game.IsGameOver() {
		return game.calculateScore(player)
	}

	if isMaximizing {
		maxEval := math.MinInt
		for move := 0; move < game.Opts.Cols; move++ {
			if !game.Place(move, player) {
				continue
			}

			next := MoveNode{ Move: move, Prev: node }
			eval := game.minimax(player, &next, depth - 1, bestMax, bestMin, false)
			if eval > maxEval {
				maxEval = eval
				node.Next = &next
			}
			if eval > bestMax {
				bestMax = eval
			}
			if bestMax >= bestMin {
				game.Unplace(move)
				break
			}

			game.Unplace(move)
		}
		return maxEval
	} else {
		minEval := math.MaxInt
		for move := 0; move < game.Opts.Cols; move++ {
			if !game.Place(move, player ^ PlayerXor) {
				continue
			}

			next := MoveNode{ Move: move, Prev: node }
			eval := game.minimax(player, &next, depth - 1, bestMax, bestMin, true)
			if eval < minEval {
				minEval = eval
				node.Next = &next
			}
			if eval < bestMin {
				bestMin = eval
			}
			if bestMin <= bestMax {
				game.Unplace(move)
				break
			}

			game.Unplace(move)
		}
		return minEval
	}
}

func (game *Game) calculateScore(player Space) int {
	// TODO: Only take into account consecutive moves that have the possibility of getting 4 in a row - take consecutive empty
	// spaces into account (change getAllConsecutives to take in a parameter which also returns consecutive empty spaces)
	score := 0

	p1Consecs, p2Consecs := game.getAllConsecutives()
	p1Multi, p2Multi := 1, -1
	p1Win, p2Win := math.MaxInt, math.MinInt
	if player == PlayerTwo {
		p1Multi, p2Multi = p2Multi, p1Multi
		p1Win, p2Win = p2Win, p1Win
	}

	for _, p1Consec := range p1Consecs {
		if p1Consec >= game.Opts.WinCond {
			return p1Win
		}
		score += p1Consec*p1Consec * p1Multi
	}
	for _, p2Consec := range p2Consecs {
		if p2Consec >= game.Opts.WinCond {
			return p2Win
		}
		score += p2Consec*p2Consec * p2Multi
	}

	return score
}


func intLog(base, n int) int {
	result := 0
	for n > 1 {
		n /= base
		result++
	}
	return result
}
