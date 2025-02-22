# Tetris Game in Go

A simple Tetris game implementation using Go and the Fyne GUI framework.

## Technology Stack

### Programming Language
- Go 1.21
  - Modern, concurrent programming language
  - Strong static typing
  - Garbage collection
  - Built-in testing support

### Core Libraries
- [Fyne](https://fyne.io/) v2.4.3
  - Cross-platform GUI framework for Go
  - Modern, material design-inspired widgets
  - Hardware-accelerated rendering
  - Native window management
  - Event handling system

### Dependencies
- fyne.io/systray v1.10.1 - System tray integration
- github.com/fsnotify/fsnotify v1.6.0 - File system notifications
- Other Fyne-related dependencies for GUI rendering and event handling

## Features

- Classic Tetris gameplay
- Next piece preview
- Score tracking
- Game over detection
- Smooth piece movement and rotation
- Modern GUI using Fyne framework

## Implementation Details

### Game Structure

The game is structured into several key components:

1. **Main Game Logic** (`main.go`)
   - Game state management
   - Board rendering
   - Score tracking
   - Game loop implementation
   - Window and UI layout

2. **Tetromino Management** (`tetromino.go`)
   - Tetromino types and shapes
   - Color definitions
   - Piece rotation logic
   - Random piece generation

3. **Keyboard Controls** (`keyboard.go`)
   - Keyboard event handling
   - Movement controls
   - Rotation controls
   - Hard drop functionality

### Game Board

- 10x20 grid (standard Tetris size)
- Each cell represented by a rectangle
- Pieces rendered using different colors
- Next piece preview displayed on the right side
- Score displayed below the next piece preview

### Controls

- Left Arrow: Move piece left
- Right Arrow: Move piece right
- Down Arrow: Move piece down
- Up Arrow: Rotate piece
- Space: Hard drop

### Scoring System

- Points awarded for clearing lines
- 100 points per line cleared
- Score updates displayed in real-time

### Technical Implementation

1. **GUI Framework**
   - Built using Fyne v2
   - Custom container layouts
   - Responsive window sizing
   - Smooth graphics rendering

2. **Game Loop**
   - Ticker-based game updates
   - Collision detection
   - Line clearing
   - Piece locking mechanism

3. **State Management**
   - Current piece tracking
   - Next piece preview
   - Board state maintenance
   - Score tracking

## Requirements

- Go 1.16 or later
- Fyne v2
- Compatible operating system (Windows, macOS, Linux)

## Installation

1. Install Go from [golang.org](https://golang.org)
2. Install Fyne dependencies:
   ```bash
   go get fyne.io/fyne/v2
   ```
3. Clone and run the game:
   ```bash
   git clone https://github.com/y16ra/tetris.git
   cd tetris
   go run .
   ```

## Future Improvements

- High score system
- Level progression
- Sound effects
- Piece hold feature
- Game pause functionality
- Settings customization
