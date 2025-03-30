package chess

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Game struct {
	Players       []Player
	Board         Board
	CurrentPlayer *Player
	CurrentTurn   int
	GameState     int // 0: ongoing, 1: check, 2: checkmate, 3: stalemate
	MoveHistory   []Move
}

type Player struct {
	Color string
}

func (g *Game) Play() (bool, error) {
	moveNum := 0

	for g.GameState == 0 {
		moveNum++
		if moveNum%2 != 0 {
			g.CurrentPlayer = &g.Players[0]
			g.PlayerMove(g.CurrentPlayer.Color)

		} else {
			g.CurrentPlayer = &g.Players[1]
			g.PlayerMove(g.CurrentPlayer.Color)

		}
	}

	return true, nil
}

func (g *Game) PlayerMove(PlayerColor string) {
	// ClearScreen()
	g.Board.Display()

	// Request player move from the user
	fmt.Printf("%s's turn:\n", PlayerColor)
	move := g.RequestMove()
	piece := g.Board.GetCell(move.From.Row, move.From.Col).Piece

	if !piece.IsValidPiece(move, g.Board, PlayerColor) {
		fmt.Println("Invalid piece")
	}

	if !piece.InValidMoves(move.To) {
		fmt.Println("Invalid move")
	}

	piece.Move(move, &g.Board)
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	default:
		cmd = exec.Command("clear") // Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Game) RequestMove() Move {

	inputMove := MoveInput()

	moveString, valid := ValidatePreInput(inputMove)

	if !valid {
		g.RequestMove()
	}
	move := MoveParser(moveString)

	return move

}

func MoveInput() string {
	// Request player move from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	inputMove, _ := reader.ReadString('\n')
	inputMove = strings.TrimSpace(inputMove)
	return inputMove
}

func MoveParser(moveString string) Move {
	// Example input: "e2e4"
	// Parse the input move into a Move struct
	// Assuming the input is always valid and in the format "e2e4"
	rows := []int{}
	cols := []int{}

	for i, char := range moveString {
		fmt.Println("char:", string(char))
		if i%2 != 0 {
			convertedRow, _ := RowToIndex(string(char))
			rows = append(rows, convertedRow)
		} else {
			convertedCol, _ := ColumnToIndex(string(char))
			cols = append(cols, convertedCol)
		}
	}

	move := Move{
		From: Position{Row: rows[0], Col: cols[0]},
		To:   Position{Row: rows[1], Col: cols[1]},
	}

	return move
}

func ValidatePreInput(moveString string) (string, bool) {

	if len(moveString) != 4 {
		fmt.Println("Length of input != 4. Please Input a move in the format e2e4")
		return moveString, false
	}

	for i, char := range moveString {
		if i%2 != 0 {
			if Contains("01234567", string(char)) == false {
				fmt.Printf("Character %d %s is not in [01234567]. Please Input a move in the format e2e4", i, string(char))
				return moveString, false
			} else if i%2 == 0 {
				if Contains("abcdefgh", string(char)) == false {
					fmt.Printf("Character %d %s is not in [abcdefgh]. Please Input a move in the format e2e4", i, string(char))
					return moveString, false
				}
			}
		}
	}
	return moveString, true
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.CurrentPlayer
}

func (g *Game) Setup() {

	WhiteRook1 := Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 0}, InGame: true, Display: "♖", ValidMoves: []Position{}}
	WhiteKnight1 := Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 1}, InGame: true, Display: "♘", ValidMoves: []Position{}}
	WhiteBishop1 := Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 2}, InGame: true, Display: "♗", ValidMoves: []Position{}}
	WhiteQueen := Piece{Name: "Queen", Color: "White", CurrentPosition: Position{Row: 0, Col: 3}, InGame: true, Display: "♕", ValidMoves: []Position{}}
	WhiteKing := Piece{Name: "King", Color: "White", CurrentPosition: Position{Row: 0, Col: 4}, InGame: true, Display: "♔", ValidMoves: []Position{}}
	WhiteBishop2 := Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 5}, InGame: true, Display: "♗", ValidMoves: []Position{}}
	WhiteKnight2 := Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 6}, InGame: true, Display: "♘", ValidMoves: []Position{}}
	WhiteRook2 := Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 7}, InGame: true, Display: "♖", ValidMoves: []Position{}}
	WhitePawn1 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 0}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn2 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 1}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn3 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 2}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn4 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 3}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn5 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 4}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn6 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 5}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn7 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 6}, InGame: true, Display: "♙", ValidMoves: []Position{}}
	WhitePawn8 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 7}, InGame: true, Display: "♙", ValidMoves: []Position{}}

	BlackRook1 := Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 0}, InGame: true, Display: "♜", ValidMoves: []Position{}}
	BlackKnight1 := Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 1}, InGame: true, Display: "♞", ValidMoves: []Position{}}
	BlackBishop1 := Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 2}, InGame: true, Display: "♝", ValidMoves: []Position{}}
	BlackQueen := Piece{Name: "Queen", Color: "Black", CurrentPosition: Position{Row: 7, Col: 3}, InGame: true, Display: "♕", ValidMoves: []Position{}}
	BlackKing := Piece{Name: "King", Color: "Black", CurrentPosition: Position{Row: 7, Col: 4}, InGame: true, Display: "♚", ValidMoves: []Position{}}
	BlackBishop2 := Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 5}, InGame: true, Display: "♝", ValidMoves: []Position{}}
	BlackKnight2 := Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 6}, InGame: true, Display: "♞", ValidMoves: []Position{}}
	BlackRook2 := Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 7}, InGame: true, Display: "♜", ValidMoves: []Position{}}
	BlackPawn1 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 0}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn2 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 1}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn3 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 2}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn4 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 3}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn5 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 4}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn6 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 5}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn7 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 6}, InGame: true, Display: "♟", ValidMoves: []Position{}}
	BlackPawn8 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 7}, InGame: true, Display: "♟", ValidMoves: []Position{}}

	cell00 := Cell{Position: Position{Row: 0, Col: 0}, Color: "White", Piece: &WhiteRook1}
	cell01 := Cell{Position: Position{Row: 0, Col: 1}, Color: "Black", Piece: &WhiteKnight1}
	cell02 := Cell{Position: Position{Row: 0, Col: 2}, Color: "White", Piece: &WhiteBishop1}
	cell03 := Cell{Position: Position{Row: 0, Col: 3}, Color: "Black", Piece: &WhiteQueen}
	cell04 := Cell{Position: Position{Row: 0, Col: 4}, Color: "White", Piece: &WhiteKing}
	cell05 := Cell{Position: Position{Row: 0, Col: 5}, Color: "Black", Piece: &WhiteBishop2}
	cell06 := Cell{Position: Position{Row: 0, Col: 6}, Color: "White", Piece: &WhiteKnight2}
	cell07 := Cell{Position: Position{Row: 0, Col: 7}, Color: "Black", Piece: &WhiteRook2}

	cell10 := Cell{Position: Position{Row: 1, Col: 0}, Color: "White", Piece: &WhitePawn1}
	cell11 := Cell{Position: Position{Row: 1, Col: 1}, Color: "Black", Piece: &WhitePawn2}
	cell12 := Cell{Position: Position{Row: 1, Col: 2}, Color: "White", Piece: &WhitePawn3}
	cell13 := Cell{Position: Position{Row: 1, Col: 3}, Color: "Black", Piece: &WhitePawn4}
	cell14 := Cell{Position: Position{Row: 1, Col: 4}, Color: "White", Piece: &WhitePawn5}
	cell15 := Cell{Position: Position{Row: 1, Col: 5}, Color: "Black", Piece: &WhitePawn6}
	cell16 := Cell{Position: Position{Row: 1, Col: 6}, Color: "White", Piece: &WhitePawn7}
	cell17 := Cell{Position: Position{Row: 1, Col: 7}, Color: "Black", Piece: &WhitePawn8}

	cell20 := Cell{Position: Position{Row: 2, Col: 0}, Color: "Black", Piece: nil}
	cell21 := Cell{Position: Position{Row: 2, Col: 1}, Color: "White", Piece: nil}
	cell22 := Cell{Position: Position{Row: 2, Col: 2}, Color: "Black", Piece: nil}
	cell23 := Cell{Position: Position{Row: 2, Col: 3}, Color: "White", Piece: nil}
	cell24 := Cell{Position: Position{Row: 2, Col: 4}, Color: "Black", Piece: nil}
	cell25 := Cell{Position: Position{Row: 2, Col: 5}, Color: "White", Piece: nil}
	cell26 := Cell{Position: Position{Row: 2, Col: 6}, Color: "Black", Piece: nil}
	cell27 := Cell{Position: Position{Row: 2, Col: 7}, Color: "White", Piece: nil}

	cell30 := Cell{Position: Position{Row: 3, Col: 0}, Color: "White", Piece: nil}
	cell31 := Cell{Position: Position{Row: 3, Col: 1}, Color: "Black", Piece: nil}
	cell32 := Cell{Position: Position{Row: 3, Col: 2}, Color: "White", Piece: nil}
	cell33 := Cell{Position: Position{Row: 3, Col: 3}, Color: "Black", Piece: nil}
	cell34 := Cell{Position: Position{Row: 3, Col: 4}, Color: "White", Piece: nil}
	cell35 := Cell{Position: Position{Row: 3, Col: 5}, Color: "Black", Piece: nil}
	cell36 := Cell{Position: Position{Row: 3, Col: 6}, Color: "White", Piece: nil}
	cell37 := Cell{Position: Position{Row: 3, Col: 7}, Color: "Black", Piece: nil}

	cell40 := Cell{Position: Position{Row: 4, Col: 0}, Color: "Black", Piece: nil}
	cell41 := Cell{Position: Position{Row: 4, Col: 1}, Color: "White", Piece: nil}
	cell42 := Cell{Position: Position{Row: 4, Col: 2}, Color: "Black", Piece: nil}
	cell43 := Cell{Position: Position{Row: 4, Col: 3}, Color: "White", Piece: nil}
	cell44 := Cell{Position: Position{Row: 4, Col: 4}, Color: "Black", Piece: nil}
	cell45 := Cell{Position: Position{Row: 4, Col: 5}, Color: "White", Piece: nil}
	cell46 := Cell{Position: Position{Row: 4, Col: 6}, Color: "Black", Piece: nil}
	cell47 := Cell{Position: Position{Row: 4, Col: 7}, Color: "White", Piece: nil}

	cell50 := Cell{Position: Position{Row: 5, Col: 0}, Color: "White", Piece: nil}
	cell51 := Cell{Position: Position{Row: 5, Col: 1}, Color: "Black", Piece: nil}
	cell52 := Cell{Position: Position{Row: 5, Col: 2}, Color: "White", Piece: nil}
	cell53 := Cell{Position: Position{Row: 5, Col: 3}, Color: "Black", Piece: nil}
	cell54 := Cell{Position: Position{Row: 5, Col: 4}, Color: "White", Piece: nil}
	cell55 := Cell{Position: Position{Row: 5, Col: 5}, Color: "Black", Piece: nil}
	cell56 := Cell{Position: Position{Row: 5, Col: 6}, Color: "White", Piece: nil}
	cell57 := Cell{Position: Position{Row: 5, Col: 7}, Color: "Black", Piece: nil}

	cell60 := Cell{Position: Position{Row: 6, Col: 0}, Color: "Black", Piece: &BlackPawn1}
	cell61 := Cell{Position: Position{Row: 6, Col: 1}, Color: "White", Piece: &BlackPawn2}
	cell62 := Cell{Position: Position{Row: 6, Col: 2}, Color: "Black", Piece: &BlackPawn3}
	cell63 := Cell{Position: Position{Row: 6, Col: 3}, Color: "White", Piece: &BlackPawn4}
	cell64 := Cell{Position: Position{Row: 6, Col: 4}, Color: "Black", Piece: &BlackPawn5}
	cell65 := Cell{Position: Position{Row: 6, Col: 5}, Color: "White", Piece: &BlackPawn6}
	cell66 := Cell{Position: Position{Row: 6, Col: 6}, Color: "Black", Piece: &BlackPawn7}
	cell67 := Cell{Position: Position{Row: 6, Col: 7}, Color: "White", Piece: &BlackPawn8}

	cell70 := Cell{Position: Position{Row: 7, Col: 0}, Color: "Black", Piece: &BlackRook1}
	cell71 := Cell{Position: Position{Row: 7, Col: 1}, Color: "White", Piece: &BlackKnight1}
	cell72 := Cell{Position: Position{Row: 7, Col: 2}, Color: "Black", Piece: &BlackBishop1}
	cell73 := Cell{Position: Position{Row: 7, Col: 3}, Color: "White", Piece: &BlackQueen}
	cell74 := Cell{Position: Position{Row: 7, Col: 4}, Color: "Black", Piece: &BlackKing}
	cell75 := Cell{Position: Position{Row: 7, Col: 5}, Color: "White", Piece: &BlackBishop2}
	cell76 := Cell{Position: Position{Row: 7, Col: 6}, Color: "Black", Piece: &BlackKnight2}
	cell77 := Cell{Position: Position{Row: 7, Col: 7}, Color: "White", Piece: &BlackRook2}

	board := Board{Cells: [][]Cell{
		{cell00, cell01, cell02, cell03, cell04, cell05, cell06, cell07},
		{cell10, cell11, cell12, cell13, cell14, cell15, cell16, cell17},
		{cell20, cell21, cell22, cell23, cell24, cell25, cell26, cell27},
		{cell30, cell31, cell32, cell33, cell34, cell35, cell36, cell37},
		{cell40, cell41, cell42, cell43, cell44, cell45, cell46, cell47},
		{cell50, cell51, cell52, cell53, cell54, cell55, cell56, cell57},
		{cell60, cell61, cell62, cell63, cell64, cell65, cell66, cell67},
		{cell70, cell71, cell72, cell73, cell74, cell75, cell76, cell77},
	}}

	PlayerWhite := Player{Color: "White"}
	PlayerBlack := Player{Color: "Black"}

	g.Players = []Player{PlayerWhite, PlayerBlack}
	g.Board = board
	g.CurrentPlayer = &PlayerWhite
	g.CurrentTurn = 1
	g.GameState = 0
	g.MoveHistory = []Move{}

	for cellRow := range g.Board.Cells {
		for cellCol := range g.Board.Cells[cellRow] {
			cell := g.Board.GetCell(cellRow, cellCol)
			if cell.Piece != nil {
				cell.Piece.SetValidMoves(&g.Board)
			}
		}

	}

}
