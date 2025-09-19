package parser_test

import (
	"testing"

	parser "github.com/cesar-engine/core/Parser"
	"github.com/cesar-engine/core/board"
)

func TestParseFEN(t *testing.T) {
	tests := []struct {
		name      string
		fen       string
		wantError bool
	}{
		{
			name:      "Starting Position",
			fen:       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			wantError: false,
		},
		{
			name:      "Empty Board",
			fen:       "8/8/8/8/8/8/8/8 w - - 0 1",
			wantError: false,
		},
		{
			name:      "Invalid Row Length",
			fen:       "rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			wantError: true,
		},
		{
			name:      "Too Few Parts",
			fen:       "8/8/8/8/8/8/8/8 w",
			wantError: true,
		},
		{
			name:      "With En Passant Square",
			fen:       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq e3 0 1",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board.Board{}
			err := parser.ParseFEN(tt.fen, b)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseFEN(%q) error = %v, wantError %v", tt.fen, err, tt.wantError)
			}

			if err == nil && !tt.wantError {
				// sanity check: board must have 8 ranks
				if len(b.Squares) != 8 {
					t.Errorf("expected 8 ranks, got %d", len(b.Squares))
				}
			}
		})
	}
}

