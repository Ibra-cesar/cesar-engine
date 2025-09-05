package board

import "fmt"

type Square struct {
	File int
	Rank int
}

func NewSquare(rank, file int) Square {
	return Square{File: file, Rank: rank}
}

func (s Square) IsValidSquare() bool {
	return s.Rank >= 0 && s.Rank <= 7 && s.File >= 0 && s.File <= 7
}

func (s Square) StdNotation() string {
	if !s.IsValidSquare() {
		return "invalid Square"
	}

	file := string(rune('a' + s.File))
	rank := fmt.Sprintf("%d", s.Rank+1)

	return file + rank
}

// algbre Notation
func StringToSquare(nota string) (Square, error) {
	if len(nota) != 2 {
		return Square{}, fmt.Errorf("Invalid Notation")
	}

	file := int(nota[0] - 'a')
	rank := int(nota[1] - '1')

	square := Square{File: file, Rank: rank}

	fmt.Printf("file: %v, rank: %v\n", square.File, square.Rank)

	if !square.IsValidSquare() {
		return Square{}, fmt.Errorf("Invalid Square: %s", nota)
	}

	return square, nil
}

func (s Square) ToIndex() int {
	return s.Rank*8 + s.File
}

func SquareFromIndex(index int) Square {
	return Square{
		Rank: index / 8,
		File: index % 8,
	}
}
