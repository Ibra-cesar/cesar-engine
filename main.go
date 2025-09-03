package main

import (
	"fmt"

	"github.com/cesar-engine/core/board"
)

func main() {
fmt.Println("=== Debugging Board Setup ===")
	
	// Test 1: Check if piece creation works
	fmt.Println("\n1. Testing piece creation:")
	rook := board.NewPiece(board.Rook, board.White)
	fmt.Printf("White Rook: Type=%d, Color=%d\n", rook.Type, rook.Color)
	
	queen := board.NewPiece(board.Queen, board.White)
	fmt.Printf("White Queen: Type=%d, Color=%d\n", queen.Type, queen.Color)
	
	king := board.NewPiece(board.King, board.White)
	fmt.Printf("White King: Type=%d, Color=%d\n", king.Type, king.Color)
	
	// Test 2: Check basic board operations
	fmt.Println("\n2. Testing basic board operations:")
	b := board.NewBoardEmptyBoard()
	
	square := board.NewSquare(0, 0) // a1
	fmt.Printf("Square a1: %s\n", square.StdNotation())
	
	// Set and get a piece
	b.SetPiece(square, rook)
	retrieved := b.GetPiece(square)
	fmt.Printf("Set Rook, Retrieved: Type=%d, Color=%d\n", retrieved.Type, retrieved.Color)
	
	// Test 3: Test starting position setup step by step
	fmt.Println("\n3. Testing starting position setup:")
	startingBoard := board.NewStartingBoard()
	
	// Check a few key squares
	testSquares := []string{"a1", "d1", "e1", "h1", "a8", "d8", "e8", "h8"}
	for _, squareStr := range testSquares {
		sq, _ := board.StringToSquare(squareStr)
				
		piece := startingBoard.GetPiece(sq)
		fmt.Printf("%s: Type=%d, Color=%d, String=%s\n", 
			squareStr, piece.Type, piece.Color, piece.GetPieceInfo())
	}
	
	// Test 4: Print entire board
	fmt.Println("\n4. Full board:")
	fmt.Println(startingBoard)
}
