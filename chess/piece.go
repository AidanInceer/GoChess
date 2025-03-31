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
	History         []Move
}

func (p *Piece) Move(move Move, b *Board) {
	for _, validMove := range p.ValidMoves {
		if validMove == move.To {
			b.Cells[move.From.Row][move.From.Col].Piece = nil
			b.Cells[move.To.Row][move.To.Col].Piece = p

			p.UpdateCurrentPosition(move.To.Row, move.To.Col)

			p.UpdateValidMoves(b)

			p.History = append(p.History, move)
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

func (p *Piece) InValidMoves(move Position) bool {
	for _, validMove := range p.ValidMoves {
		if validMove == move {
			return true
		}
	}
	return false
}

func (p *Piece) GetName() string {
	return p.Name
}

func (p *Piece) GetColor() string {
	return p.Color
}

func (p *Piece) GetCurrentPosition() Position {
	return p.CurrentPosition
}

func (p *Piece) UpdateCurrentPosition(row int, col int) {
	p.CurrentPosition.Row = row
	p.CurrentPosition.Col = col
}

func (p *Piece) UpdateValidMoves(b *Board) {

	switch p.Name {
	case "Pawn":
		p.PawnMoves(b)
	case "Knight":
		p.KnightMoves(b)
	case "Bishop":
		p.BishopMoves(b)
	case "Rook":
		p.RookMoves(b)
	case "Queen":
		p.QueenMoves(b)
	case "King":
		p.KingMoves(b)
	}
}

func (p *Piece) PawnMoves(b *Board) {
	//get current piece position

	// determine what color it is to set the +/- op
	op := 1
	if p.Color == "Black" {
		op = -1
	}

	var ValidPositions []Position

	// if white and Piece history is Empty
	if len(p.History) == 0 && p.Color == "White" {
		ValidPositions = []Position{{p.CurrentPosition.Row + op, p.CurrentPosition.Col}, {p.CurrentPosition.Row + 2*op, p.CurrentPosition.Col}}
		// If Black and Piece history is Empty
	} else if len(p.History) == 0 && p.Color == "Black" {
		ValidPositions = []Position{{p.CurrentPosition.Row + op, p.CurrentPosition.Col}, {p.CurrentPosition.Row + 2*op, p.CurrentPosition.Col}}
		// If White and has already moved
	} else if len(p.History) != 0 && p.Color == "White" {
		// If the cell in front is empty
		if b.GetCellByRelativePosition(p.CurrentPosition, op, 0).Piece == nil {
			ValidPositions = []Position{{p.CurrentPosition.Row + op, p.CurrentPosition.Col}}
			// If the cell in front is occupied by an enemy piece
		} else {
			ValidPositions = []Position{}
		}
		// If Black and has already moved
	} else if len(p.History) != 0 && p.Color == "Black" {
		if b.GetCellByRelativePosition(p.CurrentPosition, op, 0).Piece == nil {
			ValidPositions = []Position{{p.CurrentPosition.Row + op, p.CurrentPosition.Col}}
			// If the cell in front is occupied by an enemy piece
		} else {
			ValidPositions = []Position{}
		}
	}

	takePositions := []Position{{p.CurrentPosition.Row + op, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row + op, p.CurrentPosition.Col + 1}}

	for _, position := range takePositions {
		if position.IsInBounds() {
			if b.GetCellByPosition(position).Piece == nil {
			} else if b.GetCellByPosition(position).Piece.Color != p.Color {
				ValidPositions = append(ValidPositions, position)
			}
		}
	}

	// Implement En passant

	p.ValidMoves = ValidPositions
}

func (p *Piece) KnightMoves(b *Board) {

	KnightPositions := []Position{{p.CurrentPosition.Row + 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row + 2, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col - 1}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col - 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col - 2}}
	ValidPositions := []Position{}
	for _, position := range KnightPositions {
		currentPieceColour := b.GetCellByPosition(p.CurrentPosition).Piece.Color
		if position.IsInBounds() && position.CanBeOccupied(b, currentPieceColour) {
			ValidPositions = append(ValidPositions, position)
		}
	}

	p.ValidMoves = ValidPositions
}

func (p *Piece) RookMoves(b *Board) {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	Row := b.GetRow(currentCell)
	Column := b.GetCol(currentCell)

	// Convert the Rows and Column Cells to a list of Position if its not the current position
	RowPositions := p.CheckLinearMoves(b, Row, currentCell, "row")
	ColumnPositions := p.CheckLinearMoves(b, Column, currentCell, "col")

	ValidMoves := append(RowPositions, ColumnPositions...)

	p.ValidMoves = ValidMoves
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

func (p *Piece) QueenMoves(b *Board) []Position {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	Diagonals := b.CellDiagonals(currentCell)
	Column := b.GetCol(currentCell)
	Row := b.GetRow(currentCell)

	RowPositions := p.CheckLinearMoves(b, Row, currentCell, "row")
	ColumnPositions := p.CheckLinearMoves(b, Column, currentCell, "col")

	ValidPositions := append(Diagonals, ColumnPositions...)
	ValidPositions = append(ValidPositions, RowPositions...)

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

func (p *Piece) CheckLinearMoves(board *Board, cells []Cell, currCell Cell, linearType string) []Position {

	outputPosition := []Position{}
	var Curr int
	if linearType == "col" {
		Curr = currCell.Position.Row
	} else if linearType == "row" {
		Curr = currCell.Position.Col
	}

	for i := Curr + 1; i <= 7; i++ {
		if !cells[i].Position.HasPiece(board) {
			outputPosition = append(outputPosition, cells[i].Position)
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color == currCell.Piece.Color {
			break
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color != currCell.Piece.Color {
			outputPosition = append(outputPosition, cells[i].Position)
			break
		}
	}

	for i := Curr - 1; i >= 0; i-- {
		if !cells[i].Position.HasPiece(board) {
			outputPosition = append(outputPosition, cells[i].Position)
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color == currCell.Piece.Color {
			break
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color != currCell.Piece.Color {
			outputPosition = append(outputPosition, cells[i].Position)
			break
		}
	}

	return outputPosition

}
