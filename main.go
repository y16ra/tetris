package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"time"
)

const (
	boardWidth  = 10
	boardHeight = 20
	blockSize   = 30
)

type Game struct {
	window       fyne.Window
	board        [][]bool
	container    *fyne.Container
	blocks       [][]*canvas.Rectangle
	currentPiece *Tetromino
	nextPiece    *Tetromino
	nextBlocks   [][]*canvas.Rectangle
	score        int
	scoreLabel   *widget.Label
	gameOver     bool
	ticker       *time.Ticker
}

func NewGame() *Game {
	game := &Game{
		board:      make([][]bool, boardHeight),
		blocks:     make([][]*canvas.Rectangle, boardHeight),
		nextBlocks: make([][]*canvas.Rectangle, 4),
		score:      0,
		scoreLabel: widget.NewLabel("Score: 0"),
	}

	for i := range game.board {
		game.board[i] = make([]bool, boardWidth)
		game.blocks[i] = make([]*canvas.Rectangle, boardWidth)
	}

	for i := range game.nextBlocks {
		game.nextBlocks[i] = make([]*canvas.Rectangle, 4)
	}

	return game
}

func (g *Game) createBoard() *fyne.Container {
	// Main board container
	boardContainer := container.NewWithoutLayout()
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			block := canvas.NewRectangle(color.Gray{Y: 50})
			block.Resize(fyne.NewSize(float32(blockSize-1), float32(blockSize-1)))
			block.Move(fyne.NewPos(float32(x*blockSize), float32(y*blockSize)))
			boardContainer.Add(block)
			g.blocks[y][x] = block
		}
	}

	// Side panel container (next piece and score)
	sideContainer := container.NewWithoutLayout()

	// Next piece label
	nextLabel := widget.NewLabel("Next:")
	nextLabel.Move(fyne.NewPos(float32(boardWidth*blockSize+20), 10))
	sideContainer.Add(nextLabel)

	// Next piece preview
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			block := canvas.NewRectangle(color.Gray{Y: 50})
			block.Resize(fyne.NewSize(float32(blockSize-1), float32(blockSize-1)))
			block.Move(fyne.NewPos(
				float32(boardWidth*blockSize+20+x*blockSize),
				float32(40+y*blockSize),
			))
			sideContainer.Add(block)
			g.nextBlocks[y][x] = block
		}
	}

	// Score label
	g.scoreLabel.Move(fyne.NewPos(float32(boardWidth*blockSize+20), float32(40+4*blockSize+10)))
	sideContainer.Add(g.scoreLabel)

	// Main container
	mainContainer := container.NewWithoutLayout()
	mainContainer.Add(boardContainer)
	mainContainer.Add(sideContainer)
	g.container = mainContainer

	return g.container
}

func (g *Game) updateBoard() {
	// Clear board
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if g.board[y][x] {
				g.blocks[y][x].FillColor = color.Gray{Y: 200}
			} else {
				g.blocks[y][x].FillColor = color.Gray{Y: 50}
			}
			g.blocks[y][x].Refresh()
		}
	}

	// Draw current piece
	if g.currentPiece != nil {
		for y := 0; y < len(g.currentPiece.Blocks); y++ {
			for x := 0; x < len(g.currentPiece.Blocks[y]); x++ {
				if g.currentPiece.Blocks[y][x] {
					pieceY := g.currentPiece.Y + y
					pieceX := g.currentPiece.X + x
					if pieceY >= 0 && pieceY < boardHeight && pieceX >= 0 && pieceX < boardWidth {
						g.blocks[pieceY][pieceX].FillColor = g.currentPiece.Color
						g.blocks[pieceY][pieceX].Refresh()
					}
				}
			}
		}
	}

	// Update next piece preview
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			g.nextBlocks[y][x].FillColor = color.Gray{Y: 50}
			g.nextBlocks[y][x].Refresh()
		}
	}

	if g.nextPiece != nil {
		for y := 0; y < len(g.nextPiece.Blocks); y++ {
			for x := 0; x < len(g.nextPiece.Blocks[y]); x++ {
				if g.nextPiece.Blocks[y][x] {
					g.nextBlocks[y][x].FillColor = g.nextPiece.Color
					g.nextBlocks[y][x].Refresh()
				}
			}
		}
	}
}

func (g *Game) lockPiece() {
	if g.currentPiece == nil {
		return
	}

	for y := 0; y < len(g.currentPiece.Blocks); y++ {
		for x := 0; x < len(g.currentPiece.Blocks[y]); x++ {
			if g.currentPiece.Blocks[y][x] {
				boardY := g.currentPiece.Y + y
				boardX := g.currentPiece.X + x
				if boardY >= 0 && boardY < boardHeight && boardX >= 0 && boardX < boardWidth {
					g.board[boardY][boardX] = true
				}
			}
		}
	}

	g.clearLines()
	g.spawnNewPiece()
}

func (g *Game) clearLines() {
	linesCleared := 0

	for y := boardHeight - 1; y >= 0; y-- {
		filled := true
		for x := 0; x < boardWidth; x++ {
			if !g.board[y][x] {
				filled = false
				break
			}
		}

		if filled {
			linesCleared++
			// Move all lines above down
			for moveY := y; moveY > 0; moveY-- {
				for x := 0; x < boardWidth; x++ {
					g.board[moveY][x] = g.board[moveY-1][x]
				}
			}
			// Clear top line
			for x := 0; x < boardWidth; x++ {
				g.board[0][x] = false
			}
			y++ // Check the same line again as it now contains the line that was above
		}
	}

	if linesCleared > 0 {
		g.score += linesCleared * 100
		g.scoreLabel.SetText(fmt.Sprintf("Score: %d", g.score))
	}
}

func (g *Game) spawnNewPiece() {
	if g.nextPiece == nil {
		g.nextPiece = RandomTetromino()
	}
	g.currentPiece = g.nextPiece
	g.nextPiece = RandomTetromino()

	if !g.canMove(0, 0) {
		g.gameOver = true
		g.ticker.Stop()
		g.scoreLabel.SetText(fmt.Sprintf("Game Over! Final Score: %d", g.score))
	}
	g.updateBoard()
}

func (g *Game) canMove(dx, dy int) bool {
	if g.currentPiece == nil {
		return false
	}

	for y := 0; y < len(g.currentPiece.Blocks); y++ {
		for x := 0; x < len(g.currentPiece.Blocks[y]); x++ {
			if g.currentPiece.Blocks[y][x] {
				newX := g.currentPiece.X + x + dx
				newY := g.currentPiece.Y + y + dy

				if newX < 0 || newX >= boardWidth || newY < 0 || newY >= boardHeight {
					return false
				}

				if newY >= 0 && g.board[newY][newX] {
					return false
				}
			}
		}
	}

	return true
}

func (g *Game) start() {
	rand.Seed(time.Now().UnixNano())
	g.spawnNewPiece()
	g.ticker = time.NewTicker(500 * time.Millisecond)

	go func() {
		for range g.ticker.C {
			if g.gameOver {
				return
			}

			if g.canMove(0, 1) {
				g.currentPiece.Y++
				g.updateBoard()
			} else {
				g.lockPiece()
			}
		}
	}()
}

func main() {
	a := app.NewWithID("com.github.y16ra.tetris")
	window := a.NewWindow("Tetris")

	game := NewGame()
	game.window = window

	board := game.createBoard()
	window.SetContent(board)

	game.setupKeyboard()

	// Adjust window size to accommodate next piece preview
	window.Resize(fyne.NewSize(
		float32(boardWidth*blockSize+150), // Add extra width for next piece
		float32(boardHeight*blockSize+50),
	))
	window.SetFixedSize(true)
	window.Show()

	game.start()

	a.Run()
}
