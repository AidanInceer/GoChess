package chess

type Cell struct {
	Position Position
	Color    string
	Piece    *Piece
}

func (c *Cell) GetPosition() Position {
	return c.Position
}

func (c *Cell) GetPiece() *Piece {
	if c.Piece == nil {
		return nil
	}
	return c.Piece
}

func (c *Cell) GetColor() string {
	return c.Color
}
