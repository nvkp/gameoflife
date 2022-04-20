package render

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func renderSample() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	rect1 := canvas.NewRectangle(color.White)
	rect1.Resize(fyne.NewSize(100, 100))

	rect2 := canvas.NewRectangle(color.Black)
	rect2.Resize(fyne.NewSize(100, 100))

	rect3 := canvas.NewRectangle(color.Black)
	rect3.Resize(fyne.NewSize(100, 100))

	rect4 := canvas.NewRectangle(color.White)
	rect4.Resize(fyne.NewSize(100, 100))

	grid := container.New(layout.NewGridLayout(2), rect1, rect2, rect3, rect4)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 400))

	go func() {
		fmt.Println("go func zacatek")
		defer fmt.Println("go func konec")
		time.Sleep(3 * time.Second)
		newGrid := container.New(layout.NewGridLayout(2), rect2, rect1, rect4, rect3)
		myWindow.SetContent(newGrid)
	}()

	myWindow.ShowAndRun()
}
