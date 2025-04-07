package chess

type Move struct {
	From     Position
	To       Position
	MoveType string
}

func (move *Move) IsValid() bool {
	return move.To.IsInBounds()
}
