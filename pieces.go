//go:generate fyne bundle -o bundled.go pieces

package main

import "fyne.io/fyne"

func resourceForPiece() fyne.Resource {
	return resourceBlackQueenSvg
}
