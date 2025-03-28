package main

import (
	"chess/chess"
	"fmt"
)

func main() {

	game := chess.Game{}
	game.Setup()
	game.Board.Display()

	move := chess.Move{From: chess.Position{Row: 0, Col: 0}, To: chess.Position{Row: 3, Col: 4}}
	game.Board.Move(&move)
	fmt.Printf("Moved Piece '%s' from (%d,%d) to (%d,%d)\n", game.Board.GetCell(move.From.Row, move.From.Col).Piece.Name, move.From.Row, move.From.Col, move.To.Row, move.To.Col)
	game.Board.Display()

	l := game.Board.GetDiagonals(*game.Board.GetCell(3, 4))

	fmt.Println(l)

}
