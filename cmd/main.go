package main

import "github.com/lucascda/battleship-go/pkg/board"

func main() {
	b := board.NewBoard()
	b.InitializeBoardFields()
	b.RenderBoard()
}
