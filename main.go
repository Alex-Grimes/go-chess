package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	game := chess.NewGame()
	grid := createGrid(game.Position().Board())
	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))

	go func() {
		rand.Seed(time.Now().UnixNano())
		for game.Outcome() == chess.NoOutcome {
			time.Sleep(time.Microsecond * 500)
			valid := game.ValidMoves()
			m := valid[rand.Intn(len(valid))]
			move(m, game)
		}
	}()
	w.ShowAndRun()
}

func createGrid(b *chess.Board) *fyne.Container {
	grid := container.NewGridWithColumns(8)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0xE0}
			}
			p := b.Piece(chess.Square(x + (7-y)*8))
			img := canvas.NewImageFromResource(resourceForPiece(p))
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))
		}
	}

	return grid
}

func move(m chess.Move, game *chess.Game) {
	game.Move(m)
	grid := createGrid(game.Position().Board())
	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))
}
