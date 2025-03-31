package chess

import (
	"fmt"
	"strings"
)

type Board struct {
	Cells [][]Cell
}

func (board *Board) Display() {
	for i := len(board.Cells) - 1; i >= 0; i-- {
		row := board.Cells[i]
		r := []string{}
		for _, cell := range row {
			if cell.Piece != nil {
				r = append(r, " "+cell.Piece.Display+" ")
			} else {
				r = append(r, " _ ")
			}
		}
		fmt.Printf("%d %s\n", i+1, strings.Join(r, ""))
	}
	// Print column numbers
	fmt.Println("   a  b  c  d  e  f  g  h")
}

func (board *Board) GetCell(row int, col int) Cell {
	return board.Cells[row][col]
}

func (board *Board) GetCellByPosition(position Position) Cell {
	return board.Cells[position.Row][position.Col]
}

func (board *Board) GetCellByRelativePosition(position Position, row int, col int) Cell {
	_row := position.Row + row
	_col := position.Col + col

	if !board.IsInBounds(_row, _col) {
		fmt.Println(_row, _col)
		return board.Cells[_row][_col]
	}
	return board.Cells[position.Row][position.Col]
}

// boundard limit checks
func (board *Board) IsInBounds(row int, col int) bool {
	return row >= 0 && row < 8 && col >= 0 && col < 8
}

// check if the cell has a piece
func (board *Board) HasPiece(row int, col int) bool {
	return board.Cells[row][col].Piece != nil
}

func (board *Board) GetRow(currentCell Cell) []Cell {
	return board.Cells[currentCell.Position.Row]
}

func (board *Board) GetRowPositions(currentCell Cell) []Position {
	row := board.GetRow(currentCell)
	positions := []Position{}
	for _, cell := range row {
		positions = append(positions, cell.Position)
	}
	return positions
}

func (board *Board) GetCol(currentCell Cell) []Cell {
	col := make([]Cell, 8)
	for i := 0; i < 8; i++ {
		col[i] = board.Cells[i][currentCell.Position.Col]
	}
	return col
}

func (board *Board) GetColPositions(currentCell Cell) []Position {
	col := board.GetCol(currentCell)
	positions := []Position{}
	for _, cell := range col {
		positions = append(positions, cell.Position)
	}
	return positions
}
