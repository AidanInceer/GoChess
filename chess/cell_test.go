package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell_GetPosition(t *testing.T) {
	cell := Cell{Position: Position{Row: 0, Col: 0}}
	assert.Equal(t, Position{Row: 0, Col: 0}, cell.GetPosition())
}

func TestCell_GetPiece(t *testing.T) {
	cell := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	assert.Equal(t, "Pawn", cell.GetPiece().Name)
	assert.Equal(t, "White", cell.GetPiece().Color)
}

func TestCell_GetColor(t *testing.T) {
	cell := Cell{Color: "White"}
	assert.Equal(t, "White", cell.GetColor())
}

func TestCell_GetPawnPiece(t *testing.T) {
	cell := Cell{Piece: &Piece{Name: "Pawn", Color: "White"}}
	assert.Equal(t, "Pawn", cell.GetPiece().Name)
	assert.Equal(t, "White", cell.GetPiece().Color)
}

func TestCell_GetQueenPiece(t *testing.T) {
	cell := Cell{Piece: &Piece{Name: "Queen", Color: "White"}}
	assert.Equal(t, "Queen", cell.GetPiece().Name)
	assert.Equal(t, "White", cell.GetPiece().Color)
}
