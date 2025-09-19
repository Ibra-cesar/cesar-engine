package parser

import (
	"fmt"
	"strings"

	"github.com/cesar-engine/core/board"
)

const (
	StartingPositionFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNT w KQkq - 0 1"
)

func ParseFEN(fenString string, b *board.Board) error {
	parts := strings.Fields(fenString)

	if len(parts) != 6 {
		return fmt.Errorf("ERROR: INVALID FEN NOTATION, expect 6 parts but got %d", len(parts))
	}

	fenBoard := strings.Split(parts[0], "/")
	if len(fenBoard) != 8{
		return fmt.Errorf("ERROR: INVALID FEN NOTATION, incorrect notaion %v", fenBoard)
	}

	for rank, row := range fenBoard {
		file := 0
		for _, char := range row {
			if char >= '1' && char <= '8' {
				emptySquare := int(char - '0')
				for range emptySquare{
					b.Squares[rank][file] = board.Piece{Type: board.Empty, Color: board.NoColor}
					file++
				}
			}else {
				piece := board.PieceFromCharFEN(char)
				b.Squares[rank][file] = piece
				file++
			}
		}

		if file != 8 {
			return  fmt.Errorf("ERROR: INVALID FEN NOTATION, incorrect notation %v", rank)
		}
	}
	return nil
}
