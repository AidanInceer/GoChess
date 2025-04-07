package chess

import (
	"fmt"
	"sort"
	"strings"
)

type Position struct {
	Row int
	Col int
}

func (p *Position) IsInBounds() bool {
	return p.Row >= 0 && p.Row < 8 && p.Col >= 0 && p.Col < 8
}

func InBounds(p Position) bool {
	return p.Row >= 0 && p.Row < 8 && p.Col >= 0 && p.Col < 8
}

func (p *Position) IsEqual(otherPosition *Position) bool {
	return p.Row == otherPosition.Row && p.Col == otherPosition.Col
}

func (p *Position) Display() string {
	rowVal, _ := IndexToRow(p.Row)
	colVal, _ := IndexToColumn(p.Col)
	return fmt.Sprintf("%s%s", colVal, rowVal)
}

func DisplayListOfPositions(positions []Position) string {
	var output []string

	for _, pos := range positions {
		rowVal, _ := IndexToRow(pos.Row)
		colVal, _ := IndexToColumn(pos.Col)
		output = append(output, colVal+rowVal)
	}
	sort.Strings(output)
	result := strings.Join(output, ", ")

	return result
}

func DisplayPositions(Cells []Cell) string {
	var output []string

	for _, cell := range Cells {
		rowVal, _ := IndexToRow(cell.Position.Row)
		colVal, _ := IndexToColumn(cell.Position.Col)
		output = append(output, colVal+rowVal)
	}
	sort.Strings(output)
	result := "Positions: " + strings.Join(output, ", ")

	return result
}

func (p *Position) CanBeOccupied(board *Board, currentPieceColor string) bool {
	if !p.HasPiece(board) || p.IsTakeable(board, currentPieceColor) {
		return true
	}
	return false
}

func (p *Position) IsTakeable(board *Board, currentPieceColor string) bool {
	if p.HasPiece(board) && board.GetCellByPosition(*p).Piece.Color != currentPieceColor {
		return true
	}
	return false
}

func (p *Position) HasPiece(b *Board) bool {
	if b.GetCellByPosition(*p).Piece == nil {
		return false
	}
	return true
}

func IsInValidMoves(moves []Position, positions []Position) bool {
	for _, move := range moves {
		for _, position := range positions {
			if move.IsEqual(&position) {
				return true
			}
		}
	}
	return false
}
