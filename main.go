package main

import (
	"fmt"

	"github.com/cesar-engine/core/board"
)

func main() {
	fmt.Println("\n3. Testing starting position setup:")
	startingBoard := board.NewStartingBoard()

	// Check a few key squares
	testSquares := []string{"a1", "d1", "e1", "h1", "a8", "d8", "e8", "h8"}
	for _, squareStr := range testSquares {
		sq, err := board.StringToSquare(squareStr)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}

		piece := startingBoard.GetPiece(sq)
		fmt.Printf("%s: Type=%d, Color=%d, String=%s\n",
			squareStr, piece.Type, piece.Color, piece.GetPieceInfo())
	}

	// Test 4: Print entire board
	fmt.Println("\n4. Full board:")
	fmt.Println(startingBoard)
}
