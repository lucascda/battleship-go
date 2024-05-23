package board

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Board struct {
	Grid [10][10]BoardField
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) InitializeBoardFields() {

}

func (b *Board) RenderBoard() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"})
	for i, arr := range b.Grid {
		t.AppendRow([]interface{}{i + 1, arr[0].fieldChar, arr[1].fieldChar, arr[2].fieldChar, arr[3].fieldChar, arr[4].fieldChar, arr[5].fieldChar, arr[6].fieldChar, arr[7].fieldChar, arr[8].fieldChar, arr[9].fieldChar})

	}

	t.Render()
}
