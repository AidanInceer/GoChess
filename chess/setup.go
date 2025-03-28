package chess

func (g *Game) Setup() {

	WhiteRook1 := Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 0}, InGame: true, Display: "♖"}
	WhiteKnight1 := Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 1}, InGame: true, Display: "♘"}
	WhiteBishop1 := Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 2}, InGame: true, Display: "♗"}
	WhiteQueen := Piece{Name: "Queen", Color: "White", CurrentPosition: Position{Row: 0, Col: 3}, InGame: true, Display: "♕"}
	WhiteKing := Piece{Name: "King", Color: "White", CurrentPosition: Position{Row: 0, Col: 4}, InGame: true, Display: "♔"}
	WhiteBishop2 := Piece{Name: "Bishop", Color: "White", CurrentPosition: Position{Row: 0, Col: 5}, InGame: true, Display: "♗"}
	WhiteKnight2 := Piece{Name: "Knight", Color: "White", CurrentPosition: Position{Row: 0, Col: 6}, InGame: true, Display: "♘"}
	WhiteRook2 := Piece{Name: "Rook", Color: "White", CurrentPosition: Position{Row: 0, Col: 7}, InGame: true, Display: "♖"}
	WhitePawn1 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 0}, InGame: true, Display: "♙"}
	WhitePawn2 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 1}, InGame: true, Display: "♙"}
	WhitePawn3 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 2}, InGame: true, Display: "♙"}
	WhitePawn4 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 3}, InGame: true, Display: "♙"}
	WhitePawn5 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 4}, InGame: true, Display: "♙"}
	WhitePawn6 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 5}, InGame: true, Display: "♙"}
	WhitePawn7 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 6}, InGame: true, Display: "♙"}
	WhitePawn8 := Piece{Name: "Pawn", Color: "White", CurrentPosition: Position{Row: 1, Col: 7}, InGame: true, Display: "♙"}

	BlackRook1 := Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 0}, InGame: true, Display: "♜"}
	BlackKnight1 := Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 1}, InGame: true, Display: "♘"}
	BlackBishop1 := Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 2}, InGame: true, Display: "♗"}
	BlackQueen := Piece{Name: "Queen", Color: "Black", CurrentPosition: Position{Row: 7, Col: 3}, InGame: true, Display: "♕"}
	BlackKing := Piece{Name: "King", Color: "Black", CurrentPosition: Position{Row: 7, Col: 4}, InGame: true, Display: "♚"}
	BlackBishop2 := Piece{Name: "Bishop", Color: "Black", CurrentPosition: Position{Row: 7, Col: 5}, InGame: true, Display: "♗"}
	BlackKnight2 := Piece{Name: "Knight", Color: "Black", CurrentPosition: Position{Row: 7, Col: 6}, InGame: true, Display: "♘"}
	BlackRook2 := Piece{Name: "Rook", Color: "Black", CurrentPosition: Position{Row: 7, Col: 7}, InGame: true, Display: "♜"}
	BlackPawn1 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 0}, InGame: true, Display: "♟"}
	BlackPawn2 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 1}, InGame: true, Display: "♟"}
	BlackPawn3 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 2}, InGame: true, Display: "♟"}
	BlackPawn4 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 3}, InGame: true, Display: "♟"}
	BlackPawn5 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 4}, InGame: true, Display: "♟"}
	BlackPawn6 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 5}, InGame: true, Display: "♟"}
	BlackPawn7 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 6}, InGame: true, Display: "♟"}
	BlackPawn8 := Piece{Name: "Pawn", Color: "Black", CurrentPosition: Position{Row: 6, Col: 7}, InGame: true, Display: "♟"}

	cell1 := Cell{Row: 0, Col: 0, Color: "White", HasPiece: true, Piece: WhiteRook1}
	cell2 := Cell{Row: 0, Col: 1, Color: "Black", HasPiece: true, Piece: WhiteKnight1}
	cell3 := Cell{Row: 0, Col: 2, Color: "White", HasPiece: true, Piece: WhiteBishop1}
	cell4 := Cell{Row: 0, Col: 3, Color: "Black", HasPiece: true, Piece: WhiteQueen}
	cell5 := Cell{Row: 0, Col: 4, Color: "White", HasPiece: true, Piece: WhiteKing}
	cell6 := Cell{Row: 0, Col: 5, Color: "Black", HasPiece: true, Piece: WhiteBishop2}
	cell7 := Cell{Row: 0, Col: 6, Color: "White", HasPiece: true, Piece: WhiteKnight2}
	cell8 := Cell{Row: 0, Col: 7, Color: "Black", HasPiece: true, Piece: WhiteRook2}

	cell9 := Cell{Row: 1, Col: 0, Color: "White", HasPiece: true, Piece: WhitePawn1}
	cell10 := Cell{Row: 1, Col: 1, Color: "Black", HasPiece: true, Piece: WhitePawn2}
	cell11 := Cell{Row: 1, Col: 2, Color: "White", HasPiece: true, Piece: WhitePawn3}
	cell12 := Cell{Row: 1, Col: 3, Color: "Black", HasPiece: true, Piece: WhitePawn4}
	cell13 := Cell{Row: 1, Col: 4, Color: "White", HasPiece: true, Piece: WhitePawn5}
	cell14 := Cell{Row: 1, Col: 5, Color: "Black", HasPiece: true, Piece: WhitePawn6}
	cell15 := Cell{Row: 1, Col: 6, Color: "White", HasPiece: true, Piece: WhitePawn7}
	cell16 := Cell{Row: 1, Col: 7, Color: "Black", HasPiece: true, Piece: WhitePawn8}

	cell17 := Cell{Row: 2, Col: 0, Color: "Black", HasPiece: false}
	cell18 := Cell{Row: 2, Col: 1, Color: "White", HasPiece: false}
	cell19 := Cell{Row: 2, Col: 2, Color: "Black", HasPiece: false}
	cell20 := Cell{Row: 2, Col: 3, Color: "White", HasPiece: false}
	cell21 := Cell{Row: 2, Col: 4, Color: "Black", HasPiece: false}
	cell22 := Cell{Row: 2, Col: 5, Color: "White", HasPiece: false}
	cell23 := Cell{Row: 2, Col: 6, Color: "Black", HasPiece: false}
	cell24 := Cell{Row: 2, Col: 7, Color: "White", HasPiece: false}

	cell25 := Cell{Row: 3, Col: 0, Color: "White", HasPiece: false}
	cell26 := Cell{Row: 3, Col: 1, Color: "Black", HasPiece: false}
	cell27 := Cell{Row: 3, Col: 2, Color: "White", HasPiece: false}
	cell28 := Cell{Row: 3, Col: 3, Color: "Black", HasPiece: false}
	cell29 := Cell{Row: 3, Col: 4, Color: "White", HasPiece: false}
	cell30 := Cell{Row: 3, Col: 5, Color: "Black", HasPiece: false}
	cell31 := Cell{Row: 3, Col: 6, Color: "White", HasPiece: false}
	cell32 := Cell{Row: 3, Col: 7, Color: "Black", HasPiece: false}

	cell33 := Cell{Row: 4, Col: 0, Color: "Black", HasPiece: false}
	cell34 := Cell{Row: 4, Col: 1, Color: "White", HasPiece: false}
	cell35 := Cell{Row: 4, Col: 2, Color: "Black", HasPiece: false}
	cell36 := Cell{Row: 4, Col: 3, Color: "White", HasPiece: false}
	cell37 := Cell{Row: 4, Col: 4, Color: "Black", HasPiece: false}
	cell38 := Cell{Row: 4, Col: 5, Color: "White", HasPiece: false}
	cell39 := Cell{Row: 4, Col: 6, Color: "Black", HasPiece: false}
	cell40 := Cell{Row: 4, Col: 7, Color: "White", HasPiece: false}

	cell41 := Cell{Row: 5, Col: 0, Color: "White", HasPiece: false}
	cell42 := Cell{Row: 5, Col: 1, Color: "Black", HasPiece: false}
	cell43 := Cell{Row: 5, Col: 2, Color: "White", HasPiece: false}
	cell44 := Cell{Row: 5, Col: 3, Color: "Black", HasPiece: false}
	cell45 := Cell{Row: 5, Col: 4, Color: "White", HasPiece: false}
	cell46 := Cell{Row: 5, Col: 5, Color: "Black", HasPiece: false}
	cell47 := Cell{Row: 5, Col: 6, Color: "White", HasPiece: false}
	cell48 := Cell{Row: 5, Col: 7, Color: "Black", HasPiece: false}

	cell49 := Cell{Row: 6, Col: 0, Color: "Black", HasPiece: true, Piece: BlackPawn1}
	cell50 := Cell{Row: 6, Col: 1, Color: "White", HasPiece: true, Piece: BlackPawn2}
	cell51 := Cell{Row: 6, Col: 2, Color: "Black", HasPiece: true, Piece: BlackPawn3}
	cell52 := Cell{Row: 6, Col: 3, Color: "White", HasPiece: true, Piece: BlackPawn4}
	cell53 := Cell{Row: 6, Col: 4, Color: "Black", HasPiece: true, Piece: BlackPawn5}
	cell54 := Cell{Row: 6, Col: 5, Color: "White", HasPiece: true, Piece: BlackPawn6}
	cell55 := Cell{Row: 6, Col: 6, Color: "Black", HasPiece: true, Piece: BlackPawn7}
	cell56 := Cell{Row: 6, Col: 7, Color: "White", HasPiece: true, Piece: BlackPawn8}

	cell57 := Cell{Row: 7, Col: 0, Color: "Black", HasPiece: true, Piece: BlackRook1}
	cell58 := Cell{Row: 7, Col: 1, Color: "White", HasPiece: true, Piece: BlackKnight1}
	cell59 := Cell{Row: 7, Col: 2, Color: "Black", HasPiece: true, Piece: BlackBishop1}
	cell60 := Cell{Row: 7, Col: 3, Color: "White", HasPiece: true, Piece: BlackQueen}
	cell61 := Cell{Row: 7, Col: 4, Color: "Black", HasPiece: true, Piece: BlackKing}
	cell62 := Cell{Row: 7, Col: 5, Color: "White", HasPiece: true, Piece: BlackBishop2}
	cell63 := Cell{Row: 7, Col: 6, Color: "Black", HasPiece: true, Piece: BlackKnight2}
	cell64 := Cell{Row: 7, Col: 7, Color: "White", HasPiece: true, Piece: BlackRook2}

	board := Board{Cells: [][]Cell{
		{cell1, cell2, cell3, cell4, cell5, cell6, cell7, cell8},
		{cell9, cell10, cell11, cell12, cell13, cell14, cell15, cell16},
		{cell17, cell18, cell19, cell20, cell21, cell22, cell23, cell24},
		{cell25, cell26, cell27, cell28, cell29, cell30, cell31, cell32},
		{cell33, cell34, cell35, cell36, cell37, cell38, cell39, cell40},
		{cell41, cell42, cell43, cell44, cell45, cell46, cell47, cell48},
		{cell49, cell50, cell51, cell52, cell53, cell54, cell55, cell56},
		{cell57, cell58, cell59, cell60, cell61, cell62, cell63, cell64},
	}}

	PlayerWhite := Player{Color: "White"}
	PlayerBlack := Player{Color: "Black"}

	g.Players = []Player{PlayerWhite, PlayerBlack}
	g.Board = board
	g.CurrentPlayer = &PlayerWhite
	g.CurrentTurn = 1
	g.GameState = 0
	g.MoveHistory = []Move{}
}
