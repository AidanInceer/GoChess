package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var WhiteRook1 = Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 0}, InGame: true, Display: "♖", ValidMoves: []Position{}}
var WhiteKnight1 = Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 1}, InGame: true, Display: "♘", ValidMoves: []Position{}}
var WhiteBishop1 = Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 2}, InGame: true, Display: "♗", ValidMoves: []Position{}}
var WhiteQueen = Piece{Name: "Queen", Color: "White", CurrentPosition: Position{Row: 0, Col: 3}, InGame: true, Display: "♕", ValidMoves: []Position{}}
var WhiteKing = Piece{Name: "King", Color: "White", CurrentPosition: Position{Row: 0, Col: 4}, InGame: true, Display: "♔", ValidMoves: []Position{}}
var WhiteBishop2 = Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 5}, InGame: true, Display: "♗", ValidMoves: []Position{}}
var WhiteKnight2 = Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 6}, InGame: true, Display: "♘", ValidMoves: []Position{}}
var WhiteRook2 = Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 7}, InGame: true, Display: "♖", ValidMoves: []Position{}}
var WhitePawn1 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 0}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn2 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 1}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn3 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 2}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn4 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 3}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn5 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 4}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn6 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 5}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn7 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 6}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var WhitePawn8 = Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 7}, InGame: true, Display: "♙", ValidMoves: []Position{}}
var BlackRook1 = Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 0}, InGame: true, Display: "♜", ValidMoves: []Position{}}
var BlackKnight1 = Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 1}, InGame: true, Display: "♞", ValidMoves: []Position{}}
var BlackBishop1 = Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 2}, InGame: true, Display: "♝", ValidMoves: []Position{}}
var BlackQueen = Piece{Name: "Queen", Color: "Black", CurrentPosition: Position{Row: 7, Col: 3}, InGame: true, Display: "♕", ValidMoves: []Position{}}
var BlackKing = Piece{Name: "King", Color: "Black", CurrentPosition: Position{Row: 7, Col: 4}, InGame: true, Display: "♚", ValidMoves: []Position{}}
var BlackBishop2 = Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 5}, InGame: true, Display: "♝", ValidMoves: []Position{}}
var BlackKnight2 = Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 6}, InGame: true, Display: "♞", ValidMoves: []Position{}}
var BlackRook2 = Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 7}, InGame: true, Display: "♜", ValidMoves: []Position{}}
var BlackPawn1 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 0}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn2 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 1}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn3 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 2}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn4 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 3}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn5 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 4}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn6 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 5}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn7 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 6}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var BlackPawn8 = Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 7}, InGame: true, Display: "♟", ValidMoves: []Position{}}
var cell00 = Cell{Position: Position{Row: 0, Col: 0}, Color: "White", Piece: &WhiteRook1}
var cell01 = Cell{Position: Position{Row: 0, Col: 1}, Color: "Black", Piece: &WhiteKnight1}
var cell02 = Cell{Position: Position{Row: 0, Col: 2}, Color: "White", Piece: &WhiteBishop1}
var cell03 = Cell{Position: Position{Row: 0, Col: 3}, Color: "Black", Piece: &WhiteQueen}
var cell04 = Cell{Position: Position{Row: 0, Col: 4}, Color: "White", Piece: &WhiteKing}
var cell05 = Cell{Position: Position{Row: 0, Col: 5}, Color: "Black", Piece: &WhiteBishop2}
var cell06 = Cell{Position: Position{Row: 0, Col: 6}, Color: "White", Piece: &WhiteKnight2}
var cell07 = Cell{Position: Position{Row: 0, Col: 7}, Color: "Black", Piece: &WhiteRook2}
var cell10 = Cell{Position: Position{Row: 1, Col: 0}, Color: "White", Piece: &WhitePawn1}
var cell11 = Cell{Position: Position{Row: 1, Col: 1}, Color: "Black", Piece: &WhitePawn2}
var cell12 = Cell{Position: Position{Row: 1, Col: 2}, Color: "White", Piece: &WhitePawn3}
var cell13 = Cell{Position: Position{Row: 1, Col: 3}, Color: "Black", Piece: &WhitePawn4}
var cell14 = Cell{Position: Position{Row: 1, Col: 4}, Color: "White", Piece: &WhitePawn5}
var cell15 = Cell{Position: Position{Row: 1, Col: 5}, Color: "Black", Piece: &WhitePawn6}
var cell16 = Cell{Position: Position{Row: 1, Col: 6}, Color: "White", Piece: &WhitePawn7}
var cell17 = Cell{Position: Position{Row: 1, Col: 7}, Color: "Black", Piece: &WhitePawn8}
var cell20 = Cell{Position: Position{Row: 2, Col: 0}, Color: "Black", Piece: nil}
var cell21 = Cell{Position: Position{Row: 2, Col: 1}, Color: "White", Piece: nil}
var cell22 = Cell{Position: Position{Row: 2, Col: 2}, Color: "Black", Piece: nil}
var cell23 = Cell{Position: Position{Row: 2, Col: 3}, Color: "White", Piece: nil}
var cell24 = Cell{Position: Position{Row: 2, Col: 4}, Color: "Black", Piece: nil}
var cell25 = Cell{Position: Position{Row: 2, Col: 5}, Color: "White", Piece: nil}
var cell26 = Cell{Position: Position{Row: 2, Col: 6}, Color: "Black", Piece: nil}
var cell27 = Cell{Position: Position{Row: 2, Col: 7}, Color: "White", Piece: nil}
var cell30 = Cell{Position: Position{Row: 3, Col: 0}, Color: "White", Piece: nil}
var cell31 = Cell{Position: Position{Row: 3, Col: 1}, Color: "Black", Piece: nil}
var cell32 = Cell{Position: Position{Row: 3, Col: 2}, Color: "White", Piece: nil}
var cell33 = Cell{Position: Position{Row: 3, Col: 3}, Color: "Black", Piece: nil}
var cell34 = Cell{Position: Position{Row: 3, Col: 4}, Color: "White", Piece: nil}
var cell35 = Cell{Position: Position{Row: 3, Col: 5}, Color: "Black", Piece: nil}
var cell36 = Cell{Position: Position{Row: 3, Col: 6}, Color: "White", Piece: nil}
var cell37 = Cell{Position: Position{Row: 3, Col: 7}, Color: "Black", Piece: nil}
var cell40 = Cell{Position: Position{Row: 4, Col: 0}, Color: "Black", Piece: nil}
var cell41 = Cell{Position: Position{Row: 4, Col: 1}, Color: "White", Piece: nil}
var cell42 = Cell{Position: Position{Row: 4, Col: 2}, Color: "Black", Piece: nil}
var cell43 = Cell{Position: Position{Row: 4, Col: 3}, Color: "White", Piece: nil}
var cell44 = Cell{Position: Position{Row: 4, Col: 4}, Color: "Black", Piece: nil}
var cell45 = Cell{Position: Position{Row: 4, Col: 5}, Color: "White", Piece: nil}
var cell46 = Cell{Position: Position{Row: 4, Col: 6}, Color: "Black", Piece: nil}
var cell47 = Cell{Position: Position{Row: 4, Col: 7}, Color: "White", Piece: nil}
var cell50 = Cell{Position: Position{Row: 5, Col: 0}, Color: "White", Piece: nil}
var cell51 = Cell{Position: Position{Row: 5, Col: 1}, Color: "Black", Piece: nil}
var cell52 = Cell{Position: Position{Row: 5, Col: 2}, Color: "White", Piece: nil}
var cell53 = Cell{Position: Position{Row: 5, Col: 3}, Color: "Black", Piece: nil}
var cell54 = Cell{Position: Position{Row: 5, Col: 4}, Color: "White", Piece: nil}
var cell55 = Cell{Position: Position{Row: 5, Col: 5}, Color: "Black", Piece: nil}
var cell56 = Cell{Position: Position{Row: 5, Col: 6}, Color: "White", Piece: nil}
var cell57 = Cell{Position: Position{Row: 5, Col: 7}, Color: "Black", Piece: nil}
var cell60 = Cell{Position: Position{Row: 6, Col: 0}, Color: "Black", Piece: &BlackPawn1}
var cell61 = Cell{Position: Position{Row: 6, Col: 1}, Color: "White", Piece: &BlackPawn2}
var cell62 = Cell{Position: Position{Row: 6, Col: 2}, Color: "Black", Piece: &BlackPawn3}
var cell63 = Cell{Position: Position{Row: 6, Col: 3}, Color: "White", Piece: &BlackPawn4}
var cell64 = Cell{Position: Position{Row: 6, Col: 4}, Color: "Black", Piece: &BlackPawn5}
var cell65 = Cell{Position: Position{Row: 6, Col: 5}, Color: "White", Piece: &BlackPawn6}
var cell66 = Cell{Position: Position{Row: 6, Col: 6}, Color: "Black", Piece: &BlackPawn7}
var cell67 = Cell{Position: Position{Row: 6, Col: 7}, Color: "White", Piece: &BlackPawn8}
var cell70 = Cell{Position: Position{Row: 7, Col: 0}, Color: "Black", Piece: &BlackRook1}
var cell71 = Cell{Position: Position{Row: 7, Col: 1}, Color: "White", Piece: &BlackKnight1}
var cell72 = Cell{Position: Position{Row: 7, Col: 2}, Color: "Black", Piece: &BlackBishop1}
var cell73 = Cell{Position: Position{Row: 7, Col: 3}, Color: "White", Piece: &BlackQueen}
var cell74 = Cell{Position: Position{Row: 7, Col: 4}, Color: "Black", Piece: &BlackKing}
var cell75 = Cell{Position: Position{Row: 7, Col: 5}, Color: "White", Piece: &BlackBishop2}
var cell76 = Cell{Position: Position{Row: 7, Col: 6}, Color: "Black", Piece: &BlackKnight2}
var cell77 = Cell{Position: Position{Row: 7, Col: 7}, Color: "White", Piece: &BlackRook2}

var board = Board{Cells: [][]Cell{
	{cell00, cell01, cell02, cell03, cell04, cell05, cell06, cell07},
	{cell10, cell11, cell12, cell13, cell14, cell15, cell16, cell17},
	{cell20, cell21, cell22, cell23, cell24, cell25, cell26, cell27},
	{cell30, cell31, cell32, cell33, cell34, cell35, cell36, cell37},
	{cell40, cell41, cell42, cell43, cell44, cell45, cell46, cell47},
	{cell50, cell51, cell52, cell53, cell54, cell55, cell56, cell57},
	{cell60, cell61, cell62, cell63, cell64, cell65, cell66, cell67},
	{cell70, cell71, cell72, cell73, cell74, cell75, cell76, cell77},
}}

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
