package board

import (
	"fmt"

)

type Color int
type PieceType int

const (
	White Color = iota + 1
	Black
	NoColor
)

const (
	Empty PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
}

func NewPiece(pieceType PieceType, color Color) Piece {
	return Piece{Type: pieceType, Color: color}
}

func (pt PieceType) PieceTypeString() string {
	switch pt {
	case Empty:
		return "."
	case Pawn:
		return "P"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Bishop:
		return "B"
	case Knight:
		return "N"
	case Rook:
		return "R"
	default:
		return "?"
	}
}

func (c Color) PieceColorString() string {
	switch c {
		case White:
			return  "White"
		case Black:
			return  "Black"
		default:
			return "NoColor"
	}
}
func (p Piece) IsEmptyPiece() bool {
	return  p.Type == Empty
}

func (p Piece) GetPieceInfo() string {
	if p.IsEmptyPiece() {
		return  "."
	}

	pieceSymbol := p.Type.PieceTypeString()
	if p.Color == Black {
		return  fmt.Sprintf("%c", rune(pieceSymbol[0] + 32))
	}

	return pieceSymbol
}

func PieceFromCharFEN(c rune) Piece {
	var pieceType PieceType
	var color Color

	if c >= 'A' && c <= 'Z' {
		color = White
	}else if c >= 'a' && c <= 'z' {
		color = Black
		c = c - 32
	}else {
		return Piece{Type: Empty, Color: NoColor}
	}

	switch c {
	case 'P':
		pieceType = Pawn
	case 'N':
		pieceType = Knight
	case 'B':
		pieceType = Bishop
	case 'R':
		pieceType = Rook
	case 'Q':
		pieceType = Queen
	case 'K':
		pieceType = King
	default:
		return Piece{Type: Empty, Color: NoColor}
	}

	return Piece{Type: pieceType, Color: color}
}
