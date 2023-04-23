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
	if availableCols <= 1 { // just in case of log1 (infinite loop)
		availableCols = 2
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
	score := 0
	allConsecs := game.getAllConsecutives()

	p1Multi, p2Multi := 1, -1
	p1Win, p2Win := math.MaxInt, math.MinInt
	if player == PlayerTwo {
		p1Multi, p2Multi = p2Multi, p1Multi
		p1Win, p2Win = p2Win, p1Win
	}

	for _, consecs := range allConsecs {
		for i, consec := range consecs {
			// Return an "infinite" score if a player has won
			if consec.Type != Empty && consec.Count >= game.Opts.WinCond {
				switch consec.Type {
				case PlayerOne:
					return p1Win
				case PlayerTwo:
					return p2Win
				}
			}

			// Count number of consecutive empty spaces around this sequence of consecutive spaces
			// in order to calculate the maximum consecutive spaces possible
			emptySpacesAround := 0
			if i > 0 && consecs[i - 1].Type == Empty {
				emptySpacesAround += consecs[i - 1].Count
			}
			if i + 1 < len(consecs) && consecs[i + 1].Type == Empty {
				emptySpacesAround += consecs[i + 1].Count
			}

			// Only take this sequence of consecutive spaces into account if it has the
			// possibility of making it to the amount of consecutive spaces required to win
			if count := consec.Count; count + emptySpacesAround >= game.Opts.WinCond {
				switch consec.Type {
				case PlayerOne:
					score += 2*count*count * p1Multi
				case PlayerTwo:
					score += 2*count*count * p2Multi
				}
			}
		}
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
