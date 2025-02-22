package main

import (
	"image/color"
	"math/rand"
)

// Tetromino types and their shapes and colors.
type TetrominoType int

const (
	I TetrominoType = iota
	J
	L
	O
	S
	T
	Z
)

// Tetromino represents a tetromino with its type, blocks, color, and position.
type Tetromino struct {
	Blocks [][]bool
	Color  color.RGBA
	X, Y   int
}

var tetrominoes = map[TetrominoType]struct {
	blocks [][]bool
	color  color.RGBA
}{
	I: {
		blocks: [][]bool{
			{false, false, false, false},
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
		},
		color: color.RGBA{0, 255, 255, 255}, // Cyan
	},
	J: {
		blocks: [][]bool{
			{true, false, false},
			{true, true, true},
			{false, false, false},
		},
		color: color.RGBA{0, 0, 255, 255}, // Blue
	},
	L: {
		blocks: [][]bool{
			{false, false, true},
			{true, true, true},
			{false, false, false},
		},
		color: color.RGBA{255, 165, 0, 255}, // Orange
	},
	O: {
		blocks: [][]bool{
			{true, true},
			{true, true},
		},
		color: color.RGBA{255, 255, 0, 255}, // Yellow
	},
	S: {
		blocks: [][]bool{
			{false, true, true},
			{true, true, false},
			{false, false, false},
		},
		color: color.RGBA{0, 255, 0, 255}, // Green
	},
	T: {
		blocks: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, false, false},
		},
		color: color.RGBA{128, 0, 128, 255}, // Purple
	},
	Z: {
		blocks: [][]bool{
			{true, true, false},
			{false, true, true},
			{false, false, false},
		},
		color: color.RGBA{255, 0, 0, 255}, // Red
	},
}

func NewTetromino(t TetrominoType) *Tetromino {
	template := tetrominoes[t]
	blocks := make([][]bool, len(template.blocks))
	for i := range blocks {
		blocks[i] = make([]bool, len(template.blocks[i]))
		copy(blocks[i], template.blocks[i])
	}

	return &Tetromino{
		Blocks: blocks,
		Color:  template.color,
		X:      boardWidth/2 - len(blocks[0])/2,
		Y:      0,
	}
}

func RandomTetromino() *Tetromino {
	types := []TetrominoType{I, J, L, O, S, T, Z}
	return NewTetromino(types[rand.Intn(len(types))])
}

func (t *Tetromino) Rotate() {
	if len(t.Blocks) == 2 { // O piece
		return
	}

	size := len(t.Blocks)
	rotated := make([][]bool, size)
	for i := range rotated {
		rotated[i] = make([]bool, size)
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			rotated[x][size-1-y] = t.Blocks[y][x]
		}
	}

	t.Blocks = rotated
}
