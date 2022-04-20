package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/nvkp/gameoflife/board"
)

const cellSize float32 = 20

type Renderer interface {
	Render(board <-chan board.Board)
}

type WindowRenderer struct {
}

func (r WindowRenderer) Render(board <-chan board.Board) {

	gameOfLifeApp := app.New()
	window := gameOfLifeApp.NewWindow("Game Of Life")

	done := make(chan int)
	defer close(done)

	go func() {

		for {
			select {
			case <-done:
				break
			case b := <-board:

				lay := layout.NewGridLayout(int(b.XLen))
				grid := container.New(lay)

				for _, row := range b.Board {

					for _, v := range row {

						var clr color.Gray16
						if v > 0 {
							clr = color.Black
						} else {
							clr = color.White
						}
						rectangle := canvas.NewRectangle(clr)
						rectangle.Resize(fyne.NewSize(cellSize, cellSize))
						grid.Add(rectangle)
					}
				}

				window.SetContent(grid)
				window.Resize(fyne.NewSize(float32(b.XLen)*cellSize, float32(b.YLen)*cellSize))
			}
		}
	}()

	window.ShowAndRun()
}
