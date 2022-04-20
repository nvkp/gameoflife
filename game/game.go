package game

import (
	"time"

	"github.com/nvkp/gameoflife/board"
	"github.com/nvkp/gameoflife/render"
)

type CellAlive struct {
	X uint
	Y uint
}

// hra má:
// poradi kola,
// pole bunek
// zobrazovac
// rychlost kola

type Game struct {
	round     int
	board     board.Board
	renderer  render.Renderer
	roundPace time.Duration
}

// hra má metodu pro spusteni

func NewGame(xLen, yLen uint, roundPaceMillis int, renderer render.Renderer) *Game {
	// TODO game constructor
	// TODO board

	return &Game{0, *board.NewBoard(xLen, yLen), renderer, time.Duration(roundPaceMillis * int(time.Millisecond))}
}

func (g *Game) Start(cellsAlive []CellAlive) {

	boards := make(chan board.Board)
	done := make(chan int)
	defer close(done)

	for _, c := range cellsAlive {
		g.board.Board[c.Y][c.X] = 1
	}

	go func() {
		for {
			select {
			case <-done:
				break
			default:
				boards <- g.board
				g.round = g.round + 1
				g.board.NextRound()
				time.Sleep(g.roundPace)
			}
		}
	}()

	g.renderer.Render(boards)
}
