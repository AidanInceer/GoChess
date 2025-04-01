package chess

import (
	"fmt"
	"sort"
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
			p.History = append(p.History, move)

			p.UpdateValidMoves(b)
			fmt.Println("Valid Moves: ", DisplayListOfPositions(p.ValidMoves))

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
	op := 1
	if p.Color == "Black" {
		op = -1
	}

	ValidPositions := []Position{}

	// Check forward movement
	oneStep := Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col}
	twoStep := Position{p.CurrentPosition.Row + 2*op, p.CurrentPosition.Col}

	if b.GetCellByPosition(oneStep).Piece == nil {
		ValidPositions = append(ValidPositions, oneStep)

		// If the pawn hasn't moved, it can move two steps
		if len(p.History) == 0 && b.GetCellByPosition(twoStep).Piece == nil {
			ValidPositions = append(ValidPositions, twoStep)
		}
	}

	// Check diagonal captures
	for _, position := range []Position{
		{p.CurrentPosition.Row + op, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row + op, p.CurrentPosition.Col + 1},
	} {
		if position.IsInBounds() {
			if target := b.GetCellByPosition(position).Piece; target != nil && target.Color != p.Color {
				ValidPositions = append(ValidPositions, position)
			}
		}
	}

	// TODO: Implement En passant

	p.ValidMoves = ValidPositions
}

func (p *Piece) KnightMoves(b *Board) {
	moves := []Position{
		{p.CurrentPosition.Row + 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row + 2, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row - 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row + 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col - 2},
		{p.CurrentPosition.Row - 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col - 2},
	}

	ValidPositions := make([]Position, 0, len(moves)) // Preallocate capacity
	currentColor := p.Color                           // Fetch once instead of every loop iteration

	for _, pos := range moves {
		if pos.IsInBounds() && pos.CanBeOccupied(b, currentColor) {
			ValidPositions = append(ValidPositions, pos)
		}
	}

	p.ValidMoves = ValidPositions
}

func (p *Piece) RookMoves(b *Board) {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)
	rowCells := b.GetRow(currentCell)
	colCells := b.GetCol(currentCell)

	p.ValidMoves = append(
		p.CheckLinearMoves(b, rowCells, currentCell, "row"),
		p.CheckLinearMoves(b, colCells, currentCell, "col")...,
	)
}

func (p *Piece) BishopMoves(b *Board) {
	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)
	p.ValidMoves = p.CheckDiagonalMoves(b, currentCell)
}

func (p *Piece) QueenMoves(b *Board) {
	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	// Generate all possible diagonal and linear moves
	diagonals := p.CheckDiagonalMoves(b, currentCell)
	rowPositions := p.CheckLinearMoves(b, b.GetRow(currentCell), currentCell, "row")
	columnPositions := p.CheckLinearMoves(b, b.GetCol(currentCell), currentCell, "col")

	// Combine all valid positions
	p.ValidMoves = append(append(diagonals, rowPositions...), columnPositions...)
}

func (p *Piece) KingMoves(b *Board) {
	directions := []Position{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	p.ValidMoves = []Position{}

	for _, d := range directions {
		newPos := Position{p.CurrentPosition.Row + d.Row, p.CurrentPosition.Col + d.Col}
		if newPos.IsInBounds() && newPos.CanBeOccupied(b, p.Color) && !b.CheckCellAttacked(newPos, p.Color) {
			p.ValidMoves = append(p.ValidMoves, newPos)
		}
	}
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

func (p *Piece) CheckDiagonalMoves(board *Board, currCell Cell) []Position {
	diagonals := []Position{}
	currPosition := currCell.Position
	currentRow := currPosition.Row
	currentCol := currPosition.Col

	// get up left
	for i := 1; currentRow+i < 8 && currentCol-i >= 0; i++ {
		if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol - i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol - i})
		} else if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol - i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol - i})
			break
		} else {
			break
		}
	}

	// get up right
	for i := 1; currentRow+i < 8 && currentCol+i < 8; i++ {
		if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol + i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol + i})
		} else if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol + i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol + i})
			break
		} else {
			break
		}
	}
	// get down left
	for i := 1; currentRow-i >= 0 && currentCol-i >= 0; i++ {
		if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol - i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol - i})
		} else if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol - i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol - i})
			break
		} else {
			break
		}
	}
	// get down right
	for i := 1; currentRow-i >= 0 && currentCol+i < 8; i++ {
		if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol + i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol + i})
		} else if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol + i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol + i})
			break
		} else {
			break
		}
	}

	// Sort by row first, then by column
	sort.Slice(diagonals, func(i, j int) bool {
		if diagonals[i].Row != diagonals[j].Row {
			return diagonals[i].Row < diagonals[j].Row
		}
		return diagonals[i].Col < diagonals[j].Col
	})

	return diagonals
}
