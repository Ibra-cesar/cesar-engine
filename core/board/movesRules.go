package board

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

func (b *Board) GetAllPossibleMoves(fromSq Square) []Square {
	piece := b.GetPiece(fromSq)
	if piece.IsEmptyPiece() {
		return nil
	}

	switch piece.Type {
	case Pawn:
		return b.getPawnMoves(fromSq, piece.Color)
	case Knight:
		return  b.getKnightMoves(fromSq, piece.Color)
	case Bishop:
		return b.getBishopMoves(fromSq, piece.Color)
	case Rook:
		return  b.getRookMoves(fromSq, piece.Color)
	case Queen:
		return b.getQueenMoves(fromSq, piece.Color)
	case King:
		return  b.getKingMoves(fromSq, piece.Color)
	default:
		return nil
	}
}

func (b *Board) getPawnMoves(fromSq Square, color Color) []Square {
	var moves []Square
	var direction, startRank int

	if color == White {
		direction = 1
		startRank = 1
	} else {
		direction = -1
		startRank = 6
	}

	oneForwardMove := Square{Rank: fromSq.Rank + direction, File: fromSq.File}
	if oneForwardMove.IsValidSquare() && b.isSquareEmpty(oneForwardMove) {
		moves = append(moves, oneForwardMove)

		if fromSq.Rank == startRank {
			twoForwardMove := Square{Rank: fromSq.Rank + 2*direction, File: fromSq.File}
			if twoForwardMove.IsValidSquare() && b.isSquareEmpty(twoForwardMove) {
				moves = append(moves, twoForwardMove)
			}
		}
	}

	for _, fileDelta := range []int{-1, 1} {
		captureSquare := Square{Rank: fromSq.Rank + direction, File: fromSq.File + fileDelta}
		if captureSquare.IsValidSquare() {
			targetPiece := b.GetPiece(captureSquare)

			if !targetPiece.IsEmptyPiece() && targetPiece.Color != color {
				moves = append(moves, captureSquare)
			}

			if b.enPassant != nil && captureSquare == *b.enPassant {
				moves = append(moves, captureSquare)
			}
		}
	}

	return moves
}

func (b *Board) getKnightMoves(fromSq Square, color Color) []Square {
	var moves []Square

	for _, move := range KnightMoves {
		to := Square{Rank: fromSq.Rank + move.RankDelta, File: fromSq.File + move.FileDelta}
		if to.IsValidSquare() {
			targetPiece := b.GetPiece(to)
			if targetPiece.IsEmptyPiece() || targetPiece.Color != color {
				moves = append(moves, to)
			}
		}
	}
	return moves
}

func (b *Board) getBishopMoves(fromSq Square, color Color) []Square {
	return b.getSlidingMoves(fromSq, color, BishopMovesDirection)
}

func (b *Board) getRookMoves(fromSq Square, color Color) []Square {
	return b.getSlidingMoves(fromSq, color, RookMovesDirection)
}

func (b *Board) getQueenMoves(fromSq Square, color Color) []Square {
	return b.getSlidingMoves(fromSq, color, QueenMovesDirection)
}

func (b *Board) getSlidingMoves(fromSq Square, color Color, directions []Direction) []Square {
	var moves []Square

	for _, direction := range directions {
		for distance := 1; distance < 8; distance++ {
			to := Square{
				Rank: fromSq.Rank + direction.RankDelta*distance,
				File: fromSq.Rank + direction.FileDelta*distance,
			}

			if !to.IsValidSquare() {
				break
			}

			targetPiece := b.GetPiece(to)
			if targetPiece.IsEmptyPiece() {
				moves = append(moves, to)
			} else if targetPiece.Color != color {
				moves = append(moves, to)
				break
			} else {
				break
			}
		}
	}

	return moves
}

func (b *Board) getKingMoves(fromSq Square, color Color) []Square {
	var moves []Square

	for _, move := range KingMoves {
		to := Square{
			Rank: fromSq.Rank + move.RankDelta,
			File: fromSq.File + move.FileDelta,
		}

		if to.IsValidSquare() {
			targetPiece := b.GetPiece(to)

			if targetPiece.IsEmptyPiece() || targetPiece.Color != color {
				moves = append(moves, to)
			}
		}
	}
	/*
		TODO: ADD CASTLING MOVES in KING MOVES RULES
		b.getCastlinMoves
	*/
	return moves
}


