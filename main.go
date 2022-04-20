package main

import (
	"github.com/nvkp/gameoflife/game"
	"github.com/nvkp/gameoflife/render"
)

func main() {
	var renderer render.WindowRenderer
	g := game.NewGame(30, 30, 500, renderer)

	cellsAlive := []game.CellAlive{
		{6, 5}, {7, 5}, {8, 5}, {7, 4}, {8, 4}, {9, 4},
		{9, 10}, {10, 10}, {11, 10},
	}

	//{9,10}, {10,10},{11,10}
	g.Start(cellsAlive)
}
