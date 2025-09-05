package moves

type Direction struct {
	RankDelta int
	FileDelta int
}

var (
	//Orthogonal (horizontal/vertical)
	North = Direction{1, 0}
	South = Direction{-1, 0}
	East  = Direction{0, 1}
	West  = Direction{0, -1}

	NorthEast = Direction{1, 1}
	NorthWest = Direction{1, -1}
	SouthEast = Direction{-1, 1}
	SouthWest = Direction{-1, -1}

	KnightMoves = []Direction{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}

	KingMoves = []Direction{
		North, South, East, West, NorthEast, NorthWest, SouthEast, SouthWest,
	}

	RookMovesDirection   = []Direction{North, South, East, West}
	BishopMovesDirection = []Direction{NorthEast, NorthWest, SouthEast, SouthWest}
	QueenMovesDirection  = []Direction{North, South, East, West, NorthEast, NorthWest, SouthEast, SouthWest}
)
