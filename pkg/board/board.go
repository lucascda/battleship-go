package board

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Board struct {
	grid [10][10]int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) InitializeBoardFields() {
	for i := range b.grid {
		for j := range b.grid[i] {
			b.grid[i][j] = 0
		}
	}

}

func (b *Board) RenderBoard() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"})
	for i, arr := range b.grid {
		t.AppendRow([]interface{}{i + 1, arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9]})

	}

	t.Render()
}
