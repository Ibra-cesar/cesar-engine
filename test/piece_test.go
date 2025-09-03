package board_test

import (
	"testing"

	"github.com/cesar-engine/core/board"
)

func TestNewPiece(t *testing.T) {
	testCases := []struct {
		name      string
		pieceType board.PieceType
		color     board.Color
		expected  board.Piece
	}{
		{"White King", board.King, board.White, board.NewPiece(board.King, board.White)},
		{"Black Pawn", board.Pawn, board.Black, board.NewPiece(board.Pawn, board.Black)},
		{"Empty Piece No Color", board.Empty, board.White, board.NewPiece(board.Empty, board.White)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := board.NewPiece(tc.pieceType, tc.color)
			if got.Type != tc.expected.Type || got.Color != tc.expected.Color {
				t.Errorf("NewPiece(%v, %v) got {Type: %v, Color: %v}, want {Type: %v, Color: %v}",
					tc.pieceType, tc.color, got.Type, got.Color, tc.expected.Type, tc.expected.Color)
			}
		})
	}
}

func TestPieceTypeString(t *testing.T) {
	testCases := []struct {
		name      string
		pieceType board.PieceType
		expected  string
	}{
		{"Empty Type", board.Empty, "."},
		{"Pawn Type", board.Pawn, "P"},
		{"King Type", board.King, "K"},
		{"Queen Type", board.Queen, "Q"},
		{"Bishop Type", board.Bishop, "B"},
		{"Knight Type", board.Knight, "N"},
		{"Rook Type", board.Rook, "R"},
		{"Unknown Type", board.PieceType(99), "?"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.pieceType.PieceTypeString()
			if got != tc.expected {
				t.Errorf("PieceTypeString() for %v got %q, want %q", tc.pieceType, got, tc.expected)
			}
		})
	}
}

func TestPieceColorString(t *testing.T) {
	testCases := []struct {
		name     string
		color    board.Color
		expected string
	}{
		{"White Color", board.White, "White"},
		{"Black Color", board.Black, "Black"},
		{"Unknown Color", board.Color(99), "NoColor"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.color.PieceColorString()
			if got != tc.expected {
				t.Errorf("PieceColorString() for %v got %q, want %q", tc.color, got, tc.expected)
			}
		})
	}
}

func TestIsEmptyPiece(t *testing.T) {
	testCases := []struct {
		name     string
		piece    board.Piece
		expected bool
	}{
		{"Empty Piece", board.NewPiece(board.Empty, board.White), true},
		{"White King", board.NewPiece(board.King, board.White), false},
		{"Black Pawn", board.NewPiece(board.Pawn, board.Black), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.piece.IsEmptyPiece()
			if got != tc.expected {
				t.Errorf("isEmptyPiece() for %v got %t, want %t", tc.piece, got, tc.expected)
			}
		})
	}
}

func TestGetPieceInfo(t *testing.T) {
	testCases := []struct {
		name     string
		piece    board.Piece
		expected string
	}{
		{"Empty Piece Info", board.NewPiece(board.Empty, board.White), "."},
		{"White Pawn Info", board.NewPiece(board.Pawn, board.White), "P"},
		{"White King Info", board.NewPiece(board.King, board.White), "K"},
		{"White Queen Info", board.NewPiece(board.Queen, board.White), "Q"},
		{"White Bishop Info", board.NewPiece(board.Bishop, board.White), "B"},
		{"White Knight Info", board.NewPiece(board.Knight, board.White), "N"},
		{"White Rook Info", board.NewPiece(board.Rook, board.White), "R"},
		{"Black Pawn Info", board.NewPiece(board.Pawn, board.Black), "p"},
		{"Black King Info", board.NewPiece(board.King, board.Black), "k"},
		{"Black Queen Info", board.NewPiece(board.Queen, board.Black), "q"},
		{"Black Bishop Info", board.NewPiece(board.Bishop, board.Black), "b"},
		{"Black Knight Info", board.NewPiece(board.Knight, board.Black), "n"},
		{"Black Rook Info", board.NewPiece(board.Rook, board.Black), "r"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.piece.GetPieceInfo()
			if got != tc.expected {
				t.Errorf("GetPieceInfo() for %v got %q, want %q", tc.piece, got, tc.expected)
			}
		})
	}
}
