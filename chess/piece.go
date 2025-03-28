package chess

type Piece struct {
	Name            string
	Color           string
	CurrentPosition Position
	ValidMoves      []Position
	InGame          bool
	Display         string
}

func (p *Piece) GetName() string {
	return p.Name
}

func (p *Piece) InValidMoves(move Position) bool {
	for _, validMove := range p.ValidMoves {
		if validMove == move {
			return true
		}
	}
	return false
}

func (p *Piece) GetColor() string {
	return p.Color
}

func (p *Piece) GetCurrentPosition() Position {
	return p.CurrentPosition
}

func (p *Piece) GetValidMoves() []Position {
	return p.ValidMoves
}

func (p *Piece) SetValidMoves(b Board) {

	switch p.Name {
	case "Pawn":
		p.ValidMoves = p.PawnMoves(b)
	case "Knight":
		p.ValidMoves = p.KnightMoves(b)
	case "Bishop":
		p.ValidMoves = p.BishopMoves(b)
	case "Rook":
		p.ValidMoves = p.RookMoves(b)
	case "Queen":
		p.ValidMoves = p.QueenMoves(b)
	case "King":
		p.ValidMoves = p.KingMoves(b)
	}
}

func (p *Piece) PawnMoves(b Board) []Position {
	return []Position{}
}

func (p *Piece) KnightMoves(b Board) []Position {
	return []Position{}
}

func (p *Piece) BishopMoves(b Board) []Position {
	return []Position{}
}

func (p *Piece) RookMoves(b Board) []Position {
	return []Position{}
}

func (p *Piece) QueenMoves(b Board) []Position {
	return []Position{}
}

func (p *Piece) KingMoves(b Board) []Position {
	return []Position{}
}
