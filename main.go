package main

import (
	"fmt"

	parser "github.com/cesar-engine/core/Parser"
	"github.com/cesar-engine/core/board"
)

func main() {
	b := board.NewBoardEmptyBoard()

	fmt.Println(b)

	parser.ParseFEN("7k/3N2qp/b5r1/2p1Q1N1/Pp4PK/7P/1P3p2/6r1 b KQkq e3 0 1", b)

	fmt.Println(b)
}
