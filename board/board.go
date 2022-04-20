package board

// boad má dvojrozmerne pole bunek

type Board struct {
	XLen  uint
	YLen  uint
	Board [][]uint
}

func NewBoard(xLen, yLen uint) *Board {
	return &Board{xLen, yLen, newBoard(xLen, yLen)}
}

func newBoard(xLen, yLen uint) [][]uint {
	b := make([][]uint, yLen)
	for i := range b {
		b[i] = make([]uint, xLen)
	}
	return b
}

func (board *Board) numberOfNeighborsAlive(x, y uint) (count uint) {
	b := board.Board
	// upper left
	if y > 0 && x > 0 {
		count = count + b[y-1][x-1]
	}
	// upper center
	if y > 0 {
		count = count + b[y-1][x]
	}
	// upper right
	if y > 0 && x < board.XLen-1 {
		count = count + b[y-1][x+1]
	}
	// center right
	if x < board.XLen-1 {
		count = count + b[y][x+1]
	}
	// lower right
	if x < board.XLen-1 && y < board.XLen-1 {
		count = count + b[y+1][x+1]
	}
	// lower center
	if y < board.XLen-1 {
		count = count + b[y+1][x]
	}
	// lower left
	if x > 0 && y < board.XLen-1 {
		count = count + b[y+1][x-1]
	}
	// center left
	if x > 0 {
		count = count + b[y][x-1]
	}
	return count
}

// TODO func willBeAlive(x, y int)

// Každá živá buňka s méně než dvěma živými sousedy zemře.
// Každá živá buňka se dvěma nebo třemi živými sousedy zůstává žít.
// Každá živá buňka s více než třemi živými sousedy zemře.
// Každá mrtvá buňka s právě třemi živými sousedy oživne.

func (board *Board) aliveInNextRound(x, y uint) bool {

	n := board.numberOfNeighborsAlive(x, y)
	c := board.Board[y][x]
	if c > 0 {
		return n == 2 || n == 3
	} else {
		return n == 3
	}
}

func (board *Board) NextRound() {
	// calculate new board
	b := newBoard(board.XLen, board.XLen)

	for y, row := range board.Board {

		for x, _ := range row {
			shouldBeAlive := board.aliveInNextRound(uint(x), uint(y))
			if shouldBeAlive {
				b[y][x] = 1
			} else {
				b[y][x] = 0
			}

		}
	}

	board.Board = b
}
