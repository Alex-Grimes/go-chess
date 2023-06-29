package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	grid := createGrid()
	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))
	w.ShowAndRun()
}

func createGrid() *fyne.Container {
	grid := container.NewGridWithColumns(8)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if j%2 == i%2 {
				bg.FillColor = color.Gray{0xE0}
			}
			grid.Add(bg)
		}
	}

	return grid
}
