package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (g *Game) setupKeyboard() {
	if deskCanvas, ok := g.window.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {
			if g.gameOver {
				return
			}

			switch key.Name {
			case fyne.KeyLeft:
				if g.canMove(-1, 0) {
					g.currentPiece.X--
					g.updateBoard()
				}
			case fyne.KeyRight:
				if g.canMove(1, 0) {
					g.currentPiece.X++
					g.updateBoard()
				}
			case fyne.KeyDown:
				if g.canMove(0, 1) {
					g.currentPiece.Y++
					g.updateBoard()
				} else {
					g.lockPiece()
				}
			case fyne.KeyUp:
				g.currentPiece.Rotate()
				if !g.canMove(0, 0) {
					// If rotation is not possible, rotate back
					for i := 0; i < 3; i++ {
						g.currentPiece.Rotate()
					}
				}
				g.updateBoard()
			case fyne.KeySpace:
				// Hard drop
				for g.canMove(0, 1) {
					g.currentPiece.Y++
				}
				g.lockPiece()
			}
		})
	}
}
