package chess

type Move struct {
	From Position
	To   Position
}

func (move *Move) IsValid() bool {
	return move.To.IsInBounds()
}
