//go:generate fyne bundle -o bundled.go pieces

package main

import (
	"fyne.io/fyne"
	"github.com/notnil/chess"
)

func resourceForPiece(p chess.Piece) fyne.Resource {
	switch p.Color() {
	case chess.White:
		switch p.Type() {
		case chess.Pawn:
			return resourceWhitePawnSvg
		case chess.Knight:
			return resourceWhiteKnightSvg
		case chess.Bishop:
			return resourceWhiteBishopSvg
		case chess.Rook:
			return resourceWhiteRookSvg
		case chess.Queen:
			return resourceWhiteQueenSvg
		case chess.King:
			return resourceWhiteKingSvg
		}
	case chess.Black:
		switch p.Type() {
		case chess.Pawn:
			return resourceBlackPawnSvg
		case chess.Knight:
			return resourceBlackKnightSvg
		case chess.Bishop:
			return resourceBlackBishopSvg
		case chess.Rook:
			return resourceBlackRookSvg
		case chess.Queen:
			return resourceBlackQueenSvg
		case chess.King:
			return resourceBlackKingSvg

		}
	}

	return nil
}
