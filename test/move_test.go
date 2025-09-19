package board_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/cesar-engine/core/board"
)

func squarToString(sqrs []board.Square) []string {
	var strs []string

	for _, sq := range sqrs {
		strs = append(strs, sq.StdNotation())
	}
	sort.Strings(strs)
	return strs
}

func TestPawnMoves(t *testing.T) {
	newBoard := board.NewStartingBoard()

	e2, _ := board.StringToSquare("e2")
	moves := newBoard.GetAllPossibleMoves(e2)
	expect := []string{"e3", "e4"}
	actual := squarToString(moves)

	if len(actual) != len(expect) {
		t.Errorf("Expected %v moves for white pawn at e2, got %v", expect, actual)
	}

	fmt.Printf("expect: %s, actual: %s", expect[:], actual[:])

	for i, move := range expect {
		if i >= len(actual) || actual[i] != move {
			t.Errorf("Expected move %s, got %v", move, actual)
			break
		}
	}
}
