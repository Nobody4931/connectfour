package main

type Space uint8
const (
	Empty     Space = 0b00
	PlayerOne Space = 0b01
	PlayerTwo Space = 0b10
	PlayerXor Space = 0b11
)

type Options struct {
	Cols int
	Rows int
	WinCond int
}

type Game struct {
	Opts *Options
	Board [][]Space
	heights []int
}

func NewOptions() Options {
	return Options{
		Cols: 7,
		Rows: 6,
		WinCond: 4,
	}
}

func NewGame(opts *Options) Game {
	// NOTE: Rows and columns are switched in order to store data more efficiently
	board := make([][]Space, opts.Cols)
	for col := range board {
		board[col] = make([]Space, opts.Rows) // should be implicitly empty??? idk fuck golang
	}

	return Game{
		Opts: opts,
		Board: board,
		heights: make([]int, opts.Cols),
	}
}

func (game *Game) IsValidPos(col, row int) bool {
	return 0 <= col && col < game.Opts.Cols && 0 <= row && row < game.Opts.Rows
}

func (game *Game) CanPlace(col int) bool {
	return game.heights[col] < game.Opts.Rows
}

func (game *Game) Place(col int, player Space) bool {
	if !game.CanPlace(col) {
		return false
	}
	row := game.heights[col]
	game.Board[col][row] = player
	game.heights[col]++
	return true
}

func (game *Game) Unplace(col int) bool {
	if game.heights[col] == 0 {
		return false
	}
	row := game.heights[col] - 1
	game.Board[col][row] = Empty
	game.heights[col] = row
	return true
}

func (game *Game) IsGameOver() bool {
	if winner := game.Winner(); winner != Empty {
		return true
	}

	for col := 0; col < game.Opts.Cols; col++ {
		if game.CanPlace(col) {
			return false
		}
	}

	return true
}

func (game *Game) Winner() Space {
	allConsecs := game.getAllConsecutives()

	for _, consecs := range allConsecs {
		for _, consec := range consecs {
			if consec.Type != Empty && consec.Count >= game.Opts.WinCond {
				return consec.Type
			}
		}
	}

	return Empty
}


type consecSpaces struct {
	Type Space
	Count int
}

func (game *Game) getConsecutive(col, row, colOffset, rowOffset int) consecSpaces {
	space := game.Board[col][row]
	count := 1

	for game.IsValidPos(col+colOffset, row+rowOffset) && game.Board[col+colOffset][row+rowOffset] == space {
		count++
		col += colOffset
		row += rowOffset
	}

	return consecSpaces{ Type: space, Count: count }
}

func (game *Game) getConsecutives(col, row, colOffset, rowOffset int) []consecSpaces {
	consecs := make([]consecSpaces, 0)
	for game.IsValidPos(col, row) {
		consec := game.getConsecutive(col, row, colOffset, rowOffset)
		consecs = append(consecs, consec)
		col += colOffset * consec.Count
		row += rowOffset * consec.Count
	}
	return consecs
}

func (game *Game) getAllConsecutives() [][]consecSpaces {
	colConsecs := make([]consecSpaces, 0)
	rowConsecs := make([]consecSpaces, 0)
	posConsecs := make([]consecSpaces, 0)
	negConsecs := make([]consecSpaces, 0)

	// Get column consecutives
	for col := 0; col < game.Opts.Cols; col++ {
		colConsecs = append(colConsecs, game.getConsecutives(col, 0, 0, 1)...)
	}

	// Get row consecutives
	for row := 0; row < game.Opts.Rows; row++ {
		rowConsecs = append(rowConsecs, game.getConsecutives(0, row, 1, 0)...)
	}

	// Get positive-row diagonal consecutives
	posConsecs = append(posConsecs, game.getConsecutives(0, 0, 1, 1)...)
	for col := 1; col < game.Opts.Cols; col++ {
		posConsecs = append(posConsecs, game.getConsecutives(col, 0, 1, 1)...)
	}
	for row := 1; row < game.Opts.Rows; row++ {
		posConsecs = append(posConsecs, game.getConsecutives(0, row, 1, 1)...)
	}

	// Get negative-row diagonal consecutives
	lastRow := game.Opts.Rows - 1

	negConsecs = append(negConsecs, game.getConsecutives(0, lastRow, 1, -1)...)
	for col := 1; col < game.Opts.Cols; col++ {
		negConsecs = append(negConsecs, game.getConsecutives(col, lastRow, 1, -1)...)
	}
	for row := 1; row < game.Opts.Rows; row++ {
		negConsecs = append(negConsecs, game.getConsecutives(0, row, 1, -1)...)
	}

	return [][]consecSpaces{ colConsecs, rowConsecs, posConsecs, negConsecs }
}
