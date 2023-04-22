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
	p1Consecs, p2Consecs := game.getAllConsecutives()

	for _, p1Consec := range p1Consecs {
		if p1Consec >= game.Opts.WinCond {
			return PlayerOne
		}
	}
	for _, p2Consec := range p2Consecs {
		if p2Consec >= game.Opts.WinCond {
			return PlayerTwo
		}
	}

	return Empty
}


// Counts the amount of consecutive same spaces starting at (col,row) with offset (colOffset,rowOffset)
func (game *Game) countConsecutive(col, row, colOffset, rowOffset int) int {
	count := 1
	space := game.Board[col][row]

	for game.IsValidPos(col+colOffset, row+rowOffset) && game.Board[col+colOffset][row+rowOffset] == space {
		count++
		col += colOffset
		row += rowOffset
	}

	return count
}

// Gets each player's consecutive spaces throughout the board starting at (col,row) with offset (colOffset,rowOffset)
func (game *Game) getConsecutives(col, row, colOffset, rowOffset int) ([]int, []int) {
	p1Consecs := make([]int, 0)
	p2Consecs := make([]int, 0)

	for game.IsValidPos(col, row) {
		consec := game.countConsecutive(col, row, colOffset, rowOffset)
		switch game.Board[col][row] {
		case PlayerOne:
			p1Consecs = append(p1Consecs, consec)
		case PlayerTwo:
			p2Consecs = append(p2Consecs, consec)
		}
		col += colOffset * consec
		row += rowOffset * consec
	}

	return p1Consecs, p2Consecs
}

// Get each player's consecutive spaces throughout the entire board
func (game *Game) getAllConsecutives() ([]int, []int) {
	p1Consecs := make([]int, 0)
	p2Consecs := make([]int, 0)

	populate := func(col, row, colOffset, rowOffset int) {
		p1, p2 := game.getConsecutives(col, row, colOffset, rowOffset)
		p1Consecs = append(p1Consecs, p1...)
		p2Consecs = append(p2Consecs, p2...)
	}

	// Get column consecutives
	for col := 0; col < game.Opts.Cols; col++ {
		populate(col, 0, 0, 1)
	}

	// Get row consecutives
	for row := 0; row < game.Opts.Rows; row++ {
		populate(0, row, 1, 0)
	}

	// Get positive-row diagonal consecutives
	populate(0, 0, 1, 1)
	for col := 1; col < game.Opts.Cols; col++ {
		populate(col, 0, 1, 1)
	}
	for row := 1; row < game.Opts.Rows; row++ {
		populate(0, row, 1, 1)
	}

	// Get negative-row diagonal consecutives
	lastRow := game.Opts.Rows - 1
	populate(0, lastRow, 1, -1)
	for col := 1; col < game.Opts.Cols; col++ {
		populate(col, lastRow, 1, -1)
	}
	for row := 1; row < game.Opts.Rows; row++ {
		populate(0, row, 1, -1)
	}

	return p1Consecs, p2Consecs
}
