package board_test

import (
	"testing"

	"github.com/cesar-engine/core/board"
)

func TestStartingPosition(t *testing.T) {
	newBoard := board.NewStartingBoard()

	// Test white pieces
	expectedWhitePieces := map[string]board.PieceType{
		"a1": board.Rook, "b1": board.Knight, "c1": board.Bishop, "d1": board.Queen,
		"e1": board.King, "f1": board.Bishop, "g1": board.Knight, "h1": board.Rook,
	}

	for squareStr, expectedType := range expectedWhitePieces {
		square, _ := board.StringToSquare(squareStr)
		piece := newBoard.GetPiece(square)
		if piece.Type != expectedType || piece.Color != board.White {
			t.Errorf("Expected %v white piece at %s, got %v color %v", expectedType.PieceTypeString(), squareStr, piece.Type.PieceTypeString(), piece.Color.PieceColorString())
		}
	}

	// Test white pawns
	for file := 0; file < 8; file++ {
		square := board.NewSquare(1, file)
		piece := newBoard.GetPiece(square)
		if piece.Type != board.Pawn || piece.Color != board.White {
			t.Errorf("Expected white pawn at %s, got %v", square.StdNotation(), piece)
		}
	}

	// Test black pieces
	expectedBlackPieces := map[string]board.PieceType{
		"a8": board.Rook, "b8": board.Knight, "c8": board.Bishop, "d8": board.Queen,
		"e8": board.King, "f8": board.Bishop, "g8": board.Knight, "h8": board.Rook,
	}

	for squareStr, expectedType := range expectedBlackPieces {
		square, _ := board.StringToSquare(squareStr)
		piece := newBoard.GetPiece(square)
		if piece.Type != expectedType || piece.Color != board.Black {
			t.Errorf("Expected %v black piece at %s, got %v", expectedType, squareStr, piece)
		}
	}

	// Test black pawns
	for file := 0; file < 8; file++ {
		square := board.NewSquare(6, file)
		piece := newBoard.GetPiece(square)
		if piece.Type != board.Pawn || piece.Color != board.Black {
			t.Errorf("Expected black pawn at %s, got %v", square.StdNotation(), piece)
		}
	}

	// Test empty squares
	for rank := 2; rank < 6; rank++ {
		for file := 0; file < 8; file++ {
			square := board.NewSquare(rank, file)
			piece := newBoard.GetPiece(square)
			if !piece.IsEmptyPiece() {
				t.Errorf("Expected empty square at %s, got %v", square.StdNotation(), piece)
			}
		}
	}
}
