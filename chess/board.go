package chess

import (
	"fmt"
)

type Board struct {
	Cells [][]Cell
}

type Cell struct {
	Row      int
	Col      int
	Color    string
	HasPiece bool
	Piece    Piece
}

func (board *Board) Display() *Cell {
	for i := len(board.Cells) - 1; i >= 0; i-- {
		row := board.Cells[i]
		r := []string{}
		for _, cell := range row {
			if cell.HasPiece {
				r = append(r, " "+cell.Piece.Display+" ")
			} else {
				r = append(r, " _ ")
			}
		}
		fmt.Println(row[0].Row, r)
	}
	// Print column numbers
	fmt.Println("    0   1   2   3   4   5   6   7")
	return nil
}

func (board *Board) Move(move *Move) bool {
	// check if the move is valid
	if !move.IsValid() {
		return false
	}

	fR := move.From.Row
	fC := move.From.Col
	tR := move.To.Row
	tC := move.To.Col

	if !board.IsInBounds(fR, fC) || !board.IsInBounds(tR, tC) {
		return false
	}

	fromCell := board.GetCell(fR, fC)
	toCell := board.GetCell(tR, tC)

	// check if the from cell has a piece
	if !fromCell.HasPiece {
		return false
	}

	// Check if the to cell is empty or contains a piece of the other colour which is not a king
	if toCell.HasPiece {
		if toCell.Piece.Color == fromCell.Piece.Color {
			return false
		}
	}

	// check if move is in valid move for current from piece
	if !fromCell.Piece.InValidMoves(move.To) {
		return false
	}

	fromCell.HasPiece = false
	toCell.HasPiece = true
	toCell.Piece = fromCell.Piece
	fromCell.Piece = Piece{}

	return true
}

func (board *Board) GetCell(row int, col int) *Cell {
	return &board.Cells[row][col]
}

// boundard limit checks
func (board *Board) IsInBounds(row int, col int) bool {
	return row >= 0 && row < 8 && col >= 0 && col < 8
}

// check if the cell has a piece
func (board *Board) HasPiece(row int, col int) bool {
	return board.Cells[row][col].HasPiece
}

// func (board *Board) getCurrentRow(currentCell Cell) []Cell {
// 	return board.Cells[currentCell.Row]
// }

// func (board *Board) getCurrentCol(currentCell Cell) []Cell {
// 	return board.Cells[currentCell.Col]
// }

func (board *Board) GetDiagonals(currentCell Cell) [][]int {
	diagonals := []Cell{}
	currentRow := currentCell.Row
	currentCol := currentCell.Col

	// get up left
	for i := 1; currentRow+i < 8 && currentCol-i >= 0; i++ {
		diagonals = append(diagonals, board.Cells[currentRow+i][currentCol-i])
	}

	// get up right
	for i := 1; currentRow+i < 8 && currentCol+i < 8; i++ {
		diagonals = append(diagonals, board.Cells[currentRow+i][currentCol+i])
	}

	// get down left
	for i := 1; currentRow-i >= 0 && currentCol-i >= 0; i++ {
		diagonals = append(diagonals, board.Cells[currentRow-i][currentCol-i])
	}

	// get down right
	for i := 1; currentRow-i >= 0 && currentCol+i < 8; i++ {
		diagonals = append(diagonals, board.Cells[currentRow-i][currentCol+i])
	}

	d := [][]int{}
	for _, diagonal := range diagonals {
		d = append(d, []int{diagonal.Row, diagonal.Col})
	}
	return d
}
