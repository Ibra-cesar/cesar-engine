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

	return nil
}
