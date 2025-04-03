package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_Display(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}
	board.Display()
}

func TestBoard_GetCell(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}
	cell := board.GetCell(0, 0)
	assert.Equal(t, "Pawn", cell.Piece.Name)
	assert.Equal(t, "White", cell.Piece.Color)
}

func TestBoard_GetCellByPosition(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}

	cell := board.GetCellByPosition(Position{Row: 0, Col: 0})
	assert.Equal(t, "Pawn", cell.Piece.Name)
	assert.Equal(t, "White", cell.Piece.Color)
}

func TestBoard_GetCellByRelativePosition(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}

	cell := board.GetCellByRelativePosition(Position{Row: 0, Col: 0}, 1, 1)
	assert.Equal(t, "Pawn", cell.Piece.Name)
	assert.Equal(t, "White", cell.Piece.Color)
}

func TestBoard_IsInBounds(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}
	assert.True(t, board.IsInBounds(0, 0))
	assert.False(t, board.IsInBounds(9, 9))
}

func TestBoard_HasPiece(t *testing.T) {
	board := Board{Cells: [][]Cell{
		{Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}},
		{Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}},
	}}
	assert.True(t, board.HasPiece(0, 0))
}

func TestBoard_GetRow(t *testing.T) {
	cell00 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell01 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	board := Board{Cells: [][]Cell{
		{cell00, cell01},
		{cell00, cell01},
	}}
	row := board.GetRow(cell00)
	assert.Equal(t, "Pawn", row[0].Piece.Name)
	assert.Equal(t, "White", row[0].Piece.Color)
}

func TestBoard_GetRowPositions(t *testing.T) {
	cell1 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell2 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	board := Board{Cells: [][]Cell{
		{cell1, cell2},
		{cell1, cell2},
	}}

	positions := board.GetRowPositions(cell1)
	assert.Equal(t, Position{Row: 0, Col: 0}, positions[0])
}

func TestBoard_GetCol(t *testing.T) {
	cell1 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell2 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell3 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell4 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell5 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell6 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell7 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell8 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	board := Board{Cells: [][]Cell{
		{cell1},
		{cell2},
		{cell3},
		{cell4},
		{cell5},
		{cell6},
		{cell7},
		{cell8},
	}}
	col := board.GetCol(cell1)
	assert.Equal(t, "Pawn", col[0].Piece.Name)
	assert.Equal(t, "White", col[0].Piece.Color)
	assert.Equal(t, "Pawn", col[1].Piece.Name)
	assert.Equal(t, "Black", col[1].Piece.Color)
	assert.Equal(t, "Pawn", col[2].Piece.Name)
	assert.Equal(t, "White", col[2].Piece.Color)
	assert.Equal(t, "Pawn", col[3].Piece.Name)
	assert.Equal(t, "Black", col[3].Piece.Color)

}

func TestBoard_GetColPositions(t *testing.T) {
	cell1 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell2 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell3 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell4 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell5 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell6 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	cell7 := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	cell8 := Cell{Piece: &Piece{Name: "Pawn", Color: "Black"}}
	board := Board{Cells: [][]Cell{
		{cell1},
		{cell2},
		{cell3},
		{cell4},
		{cell5},
		{cell6},
		{cell7},
		{cell8},
	}}

	positions := board.GetColPositions(cell1)
	assert.Equal(t, Position{Row: 0, Col: 0}, positions[0])

}

func TestBoard_CheckCellAttacked(t *testing.T) {
	piece1 := Piece{Name: "Pawn", Color: "White", ValidMoves: []Position{Position{Row: 1, Col: 1}, Position{Row: 2, Col: 1}}}
	cell1 := Cell{Position: Position{Row: 0, Col: 0}}
	cell2 := Cell{Position: Position{Row: 0, Col: 1}, Piece: &piece1}
	cell3 := Cell{Position: Position{Row: 0, Col: 2}}
	cell4 := Cell{Position: Position{Row: 1, Col: 0}}
	cell5 := Cell{Position: Position{Row: 1, Col: 1}}
	cell6 := Cell{Position: Position{Row: 1, Col: 2}}
	cell7 := Cell{Position: Position{Row: 2, Col: 0}}
	cell8 := Cell{Position: Position{Row: 2, Col: 1}}
	cell9 := Cell{Position: Position{Row: 2, Col: 2}}
	board := Board{Cells: [][]Cell{
		{cell1, cell2, cell3},
		{cell4, cell5, cell6},
		{cell7, cell8, cell9},
	}}
	color := "Black"
	attacked := board.CheckCellAttacked(Position{Row: 1, Col: 1}, color)
	assert.True(t, attacked)
}
