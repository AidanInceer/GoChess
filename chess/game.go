package chess

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
	return false, nil
}

func (g *Game) Move(move Move) bool {
	return g.Board.Move(&move)
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.CurrentPlayer
}
