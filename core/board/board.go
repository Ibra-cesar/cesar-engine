package board

import (
	"fmt"
	"strings"
)

type CastlingRight struct {
	WhiteKingSide  bool
	WhiteQueenSide bool
	BlackKingSide  bool
	BlackQueenSide bool
}

type Board struct {
	squares       [8][8]Piece
	castle        CastlingRight
	activeColor   Color
	enPassant     *Square
	halfmoveClock int
	fullmoveNum   int
}

func NewBoardEmptyBoard() *Board {
	board := &Board{
		squares:       [8][8]Piece{},
		castle:        CastlingRight{true, true, true, true},
		activeColor:   White,
		enPassant:     nil,
		halfmoveClock: 0,
		fullmoveNum:   1,
	}

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			board.squares[rank][file] = Piece{Type: Empty, Color: NoColor}
		}
	}

	return board
}

func NewStartingBoard() *Board {
	newBoard := NewBoardEmptyBoard()

	newBoard.squares[0][0] = NewPiece(Rook, White)
	newBoard.squares[0][1] = NewPiece(Knight, White)
	newBoard.squares[0][2] = NewPiece(Bishop, White)
	newBoard.squares[0][3] = NewPiece(Queen, White)
	newBoard.squares[0][4] = NewPiece(King, White)
	newBoard.squares[0][5] = NewPiece(Bishop, White)
	newBoard.squares[0][6] = NewPiece(Knight, White)
	newBoard.squares[0][7] = NewPiece(Rook, White)

	for file := 0; file < 8; file++ {
		newBoard.squares[1][file] = NewPiece(Pawn, White)
	}

	newBoard.squares[7][0] = NewPiece(Rook, Black)
	newBoard.squares[7][1] = NewPiece(Knight, Black)
	newBoard.squares[7][2] = NewPiece(Bishop, Black)
	newBoard.squares[7][3] = NewPiece(Queen, Black)
	newBoard.squares[7][4] = NewPiece(King, Black)
	newBoard.squares[7][5] = NewPiece(Bishop, Black)
	newBoard.squares[7][6] = NewPiece(Knight, Black)
	newBoard.squares[7][7] = NewPiece(Rook, Black)

	for file := 0; file < 8; file++ {
		newBoard.squares[6][file] = NewPiece(Pawn, Black)
	}

	newBoard.castle = CastlingRight{true, true, true, true}
	newBoard.activeColor = White
	newBoard.enPassant = nil
	newBoard.halfmoveClock = 0
	newBoard.fullmoveNum = 1

	return newBoard
}

func (b *Board) GetPiece(square Square) Piece {
	if !square.IsValidSquare() {
		return Piece{Type: Empty, Color: NoColor}
	}
	return b.squares[square.Rank][square.File]
}

func (b *Board) SetPiece(square Square, piece Piece) {
	if square.IsValidSquare() {
		b.squares[square.Rank][square.File] = piece
	}
}


func (b *Board) String() string {
	var sb strings.Builder
	
	// Print board from rank 8 to rank 1 (top to bottom)
	for rank := 7; rank >= 0; rank-- {
		sb.WriteString(fmt.Sprintf("%d ", rank+1))
		for file := 0; file < 8; file++ {
			piece := b.squares[rank][file]
			sb.WriteString(piece.GetPieceInfo() + " ")
		}
		sb.WriteString("\n")
	}
	
	// Print file letters
	sb.WriteString("  a b c d e f g h\n")
	
	// Print game state
	sb.WriteString(fmt.Sprintf("Active: %s\n", b.activeColor.PieceColorString()))
	sb.WriteString(fmt.Sprintf("Castling: K%t Q%t k%t q%t\n", 
		b.castle.WhiteKingSide, b.castle.WhiteQueenSide,
		b.castle.BlackKingSide, b.castle.BlackQueenSide))
	
	if b.enPassant != nil {
		sb.WriteString(fmt.Sprintf("En passant: %v\n", b.enPassant.StdNotation()))
	}
	
	sb.WriteString(fmt.Sprintf("Halfmove: %d, Fullmove: %d\n", b.halfmoveClock, b.fullmoveNum))
	
	return sb.String()
}

// FindKing finds the king of the specified color and returns its square
func (b *Board) FindKing(color Color) (Square, bool) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			piece := b.squares[rank][file]
			if piece.Type == King && piece.Color == color {
				return NewSquare(rank, file), true
			}
		}
	}
	return Square{}, false
}
