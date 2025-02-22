package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (g *Game) getGhostPieceY() int {
	if g.currentPiece == nil {
		return 0
	}

	ghostY := g.currentPiece.Y
	for g.canMove(0, 1) {
		g.currentPiece.Y++
	}
	finalY := g.currentPiece.Y
	g.currentPiece.Y = ghostY
	return finalY
}

func (g *Game) setupKeyboard() {
	if deskCanvas, ok := g.window.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {
			switch key.Name {
			case fyne.KeyP:
				if !g.gameOver {
					g.togglePause()
				}
				return
			case fyne.KeyR:
				g.restart()
				return
			}

			if g.gameOver || g.paused {
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
				g.currentPiece.Y = g.getGhostPieceY()
				g.lockPiece()
			}
		})
	}
}
