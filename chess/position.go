package chess

type Position struct {
	Row int
	Col int
}

func (position *Position) IsInBounds() bool {
	return position.Row >= 0 && position.Row < 8 && position.Col >= 0 && position.Col < 8
}

func (position *Position) IsEqual(otherPosition *Position) bool {
	return position.Row == otherPosition.Row && position.Col == otherPosition.Col
}
