package board

import "fmt"

type Board struct {
	grid [10][10]int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) PrintBoard() {
	letters := "ABCDEFGHIJ"
	fmt.Printf(" ")
	for _, letter := range letters {
		fmt.Printf(" %c |", letter)
	}
	fmt.Println()
	for i, _ := range letters {
		if i == 0 {
			continue
		}
		fmt.Printf("%d \n", i)
	}
	fmt.Printf("%d", 10)
}
