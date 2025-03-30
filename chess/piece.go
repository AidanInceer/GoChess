package chess

import (
	"fmt"
)

type Piece struct {
	Name            string
	Color           string
	CurrentPosition Position
	ValidMoves      []Position
	InGame          bool
	Display         string
	History         []Position
}

func (p *Piece) Move(move Move, b *Board) {
	for _, validMove := range p.ValidMoves {
		if validMove == move.To {
			b.Cells[move.From.Row][move.From.Col].Piece = nil
			b.Cells[move.To.Row][move.To.Col].Piece = p

			p.History = append(p.History, move.From)
			return
		}
	}

}

func (p *Piece) StringPosition() string {
	return fmt.Sprintf("(%d, %d)", p.CurrentPosition.Row, p.CurrentPosition.Col)
}

func (p *Piece) IsValidPiece(move Move, b Board, color string) bool {

	if p.Color != color {
		fmt.Println("That's not your piece")
		return false
	}

	if p == nil {
		fmt.Println("No piece at that location")
		return false
	}

	return true
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

func (p *Piece) GetHistory() []Position {
	return p.History
}

func (p *Piece) SetValidMoves(b *Board) {

	p.ValidMoves = []Position{}

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

func (p *Piece) PawnMoves(b *Board) []Position {
	//get current piece position

	// determine what color it is to set the +/- operator
	operator := 1
	if p.Color == "Black" {
		operator = -1
	}

	var ValidPositions []Position

	// if white and Piece history is Empty
	if len(p.History) == 0 && p.Color == "White" {
		ValidPositions = []Position{{p.CurrentPosition.Row + operator, p.CurrentPosition.Col}, {p.CurrentPosition.Row + 2*operator, p.CurrentPosition.Col}}
	} else if len(p.History) == 0 && p.Color == "Black" {
		ValidPositions = []Position{{p.CurrentPosition.Row + operator, p.CurrentPosition.Col}, {p.CurrentPosition.Row + 2*operator, p.CurrentPosition.Col}}
	} else if len(p.History) != 0 && p.Color == "White" {
		ValidPositions = []Position{{p.CurrentPosition.Row + operator, p.CurrentPosition.Col}}
	} else if len(p.History) != 0 && p.Color == "Black" {
		ValidPositions = []Position{{p.CurrentPosition.Row + operator, p.CurrentPosition.Col}}
	}

	p.ValidMoves = ValidPositions

	return ValidPositions
}

func (p *Piece) KnightMoves(b *Board) []Position {

	KnightPositions := []Position{{p.CurrentPosition.Row + 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row + 2, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col - 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col - 2}}

	ValidPositions := []Position{}
	for _, position := range KnightPositions {
		if position.IsInBounds() {
			ValidPositions = append(ValidPositions, position)
		}
	}

	p.ValidMoves = ValidPositions
	return ValidPositions

}

func (p *Piece) BishopMoves(b *Board) []Position {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	Diagonals := b.CellDiagonals(currentCell)

	ValidPositions := []Position{}

	for _, position := range Diagonals {
		if position.IsInBounds() {
			ValidPositions = append(ValidPositions, position)
		}
	}

	p.ValidMoves = ValidPositions

	return ValidPositions
}

func (p *Piece) RookMoves(b *Board) []Position {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	Columns := b.GetCol(currentCell)
	Rows := b.GetRow(currentCell)

	// Convert the Rows and Column Cells to a list of Position if its not the current position
	RowPositions := []Position{}
	ColumnPositions := []Position{}

	for _, row := range Rows {
		if row.Position != p.CurrentPosition {
			RowPositions = append(RowPositions, row.Position)
		}
	}

	for _, column := range Columns {
		if column.Position != p.CurrentPosition {
			ColumnPositions = append(ColumnPositions, column.Position)
		}
	}

	ValidPositions := append(RowPositions, ColumnPositions...)

	p.ValidMoves = ValidPositions

	return ValidPositions
}

func (p *Piece) QueenMoves(b *Board) []Position {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	Diagonals := b.CellDiagonals(currentCell)
	Columns := b.GetColPositions(currentCell)
	Rows := b.GetRowPositions(currentCell)

	ValidPositions := append(Diagonals, Columns...)
	ValidPositions = append(ValidPositions, Rows...)

	p.ValidMoves = ValidPositions

	return ValidPositions
}

func (p *Piece) KingMoves(b *Board) []Position {

	ValidPositions := []Position{}

	// get all the positions around the king
	KingPositions := []Position{{p.CurrentPosition.Row + 1, p.CurrentPosition.Col}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col}, {p.CurrentPosition.Row, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col - 1}}

	for _, position := range KingPositions {
		if position.IsInBounds() {
			ValidPositions = append(ValidPositions, position)
		}
	}

	p.ValidMoves = ValidPositions

	return ValidPositions
}
