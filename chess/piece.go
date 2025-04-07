package chess

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Piece struct {
	Name            string
	Color           string
	CurrentPosition Position
	ValidMoves      []Position
	InGame          bool
	Display         string
	History         []Move
}

func (p *Piece) Move(move Move, b *Board, moveHistory *[]MoveHistory) {
	for _, validMove := range p.ValidMoves {
		if validMove == move.To {

			p.EnPassant(move, b, moveHistory)
			b.Cells[move.From.Row][move.From.Col].Piece = nil
			b.Cells[move.To.Row][move.To.Col].Piece = p

			p.UpdateCurrentPosition(move.To.Row, move.To.Col)
			p.History = append(p.History, move)

			p.UpdateValidMoves(b, *moveHistory)
			fmt.Println("Valid Moves: ", DisplayListOfPositions(p.ValidMoves))

			*moveHistory = append(*moveHistory, MoveHistory{
				MoveNum:    len(*moveHistory) + 1,
				Move:       move,
				PieceName:  p.Name,
				PieceColor: p.Color,
			})
		}
	}

}

func (p *Piece) StringPosition() string {
	return fmt.Sprintf("(%d, %d)", p.CurrentPosition.Row, p.CurrentPosition.Col)
}

func (p *Piece) EnPassant(move Move, b *Board, moveHistory *[]MoveHistory) {
	if p.Color == "White" {
		if move.From.Row == 4 && move.To.Row == 5 && (move.From.Col == move.To.Col+1 || move.From.Col == move.To.Col-1) && b.GetCellByPosition(move.To).Piece == nil {
			// If the current move is Enpassant then remove the piece which was taken
			b.Cells[move.From.Row][move.From.Col-1].Piece.Display = "_"
			b.Cells[move.To.Row][move.From.Col-1].Piece = nil
		}
	} else {
		if move.From.Row == 3 && move.To.Row == 2 && (move.From.Col == move.To.Col+1 || move.From.Col == move.To.Col-1) && b.GetCellByPosition(move.To).Piece == nil {
			// If the current move is Enpassant then remove the piece which was taken
			b.Cells[move.To.Row][move.From.Col+1].Piece = nil
			b.Cells[move.From.Row][move.From.Col+1].Piece.Display = "_"
		}
	}
}

func (p *Piece) IsValidPiece(move Move, b Board, color string) bool {

	if p.Color != color {
		fmt.Println("That's not your piece")
		return false
	}

	if p == nil {
		fmt.Println("No piece at that location")
		return false
	}

	return true
}

func (p *Piece) InValidMoves(move Position) bool {
	for _, validMove := range p.ValidMoves {
		if validMove == move {
			return true
		}
	}
	return false
}

func (p *Piece) GetName() string {
	return p.Name
}

func (p *Piece) GetColor() string {
	return p.Color
}

func (p *Piece) GetCurrentPosition() Position {
	return p.CurrentPosition
}

func (p *Piece) UpdateCurrentPosition(row int, col int) {
	p.CurrentPosition.Row = row
	p.CurrentPosition.Col = col
}

func (p *Piece) UpdateValidMoves(b *Board, moveHistory []MoveHistory) {

	switch p.Name {
	case "Pawn":
		p.PawnMoves(b, moveHistory)
	case "Knight":
		p.KnightMoves(b)
	case "Bishop":
		p.BishopMoves(b)
	case "Rook":
		p.RookMoves(b)
	case "Queen":
		p.QueenMoves(b)
	case "King":
		p.KingMoves(b)
	}
}

func (p *Piece) PawnMoves(b *Board, moveHistory []MoveHistory) {
	op := 1
	if p.Color == "Black" {
		op = -1
	}

	ValidPositions := []Position{}

	// Check forward movement
	oneStep := Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col}
	twoStep := Position{p.CurrentPosition.Row + 2*op, p.CurrentPosition.Col}

	// Check if the one step is in bounds and can be occupied
	if oneStep.IsInBounds() && oneStep.CanBeOccupied(b, p.Color) {
		ValidPositions = append(ValidPositions, oneStep)

		// If the pawn hasn't moved, it can move two steps
		if len(p.History) == 0 && b.GetCellByPosition(twoStep).Piece == nil {
			ValidPositions = append(ValidPositions, twoStep)
		}
	}

	// Check diagonal captures
	for _, position := range []Position{
		{p.CurrentPosition.Row + op, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row + op, p.CurrentPosition.Col + 1},
	} {
		if position.IsInBounds() {
			if target := b.GetCellByPosition(position).Piece; target != nil && target.Color != p.Color {
				ValidPositions = append(ValidPositions, position)
			}
		}
	}

	// TODO: Implement En passant

	if len(moveHistory) > 0 {

		lastMove := moveHistory[len(moveHistory)-1]

		if lastMove.PieceName == "Pawn" && lastMove.PieceColor == "White" {
			if InBounds(Position{p.CurrentPosition.Row, p.CurrentPosition.Col + 1}) {
				if lastMove.Move.From.Row == 1 && lastMove.Move.To.Row == 3 && lastMove.Move.From.Col == p.CurrentPosition.Col+1 {
					ValidPositions = append(ValidPositions, Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col + 1})
				}
			}

			if InBounds(Position{p.CurrentPosition.Row, p.CurrentPosition.Col - 1}) {
				if lastMove.Move.From.Row == 1 && lastMove.Move.To.Row == 3 && lastMove.Move.From.Col == p.CurrentPosition.Col-1 {
					ValidPositions = append(ValidPositions, Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col - 1})
				}
			}
		}

		if lastMove.PieceName == "Pawn" && lastMove.PieceColor == "Black" {
			if InBounds(Position{p.CurrentPosition.Row, p.CurrentPosition.Col + 1}) {
				if lastMove.Move.From.Row == 6 && lastMove.Move.To.Row == 4 && lastMove.Move.From.Col == p.CurrentPosition.Col+1 {
					ValidPositions = append(ValidPositions, Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col + 1})

				}
			}
			if InBounds(Position{p.CurrentPosition.Row, p.CurrentPosition.Col - 1}) {
				if lastMove.Move.From.Row == 6 && lastMove.Move.To.Row == 4 && lastMove.Move.From.Col == p.CurrentPosition.Col-1 {
					ValidPositions = append(ValidPositions, Position{p.CurrentPosition.Row + op, p.CurrentPosition.Col - 1})
				}
			}
		}
	}

	// TODO: Implement Full Promotion
	p.CheckPromotion(b, moveHistory)

	p.ValidMoves = ValidPositions

}

func (p *Piece) CheckPromotion(b *Board, moveHistory []MoveHistory) {
	if p.Color == "White" {
		if p.CurrentPosition.Row == 7 {
			fmt.Println("Valid Moves: ", DisplayListOfPositions(p.ValidMoves))
			p.Name, p.Display = p.SelectPromotionPiece()
			fmt.Println("Promoting to", p.Name)
			p.UpdateValidMoves(b, moveHistory)
		}
	} else {
		if p.CurrentPosition.Row == 0 {
			fmt.Println("Valid Moves: ", DisplayListOfPositions(p.ValidMoves))
			p.Name, p.Display = p.SelectPromotionPiece()
			fmt.Println("Promoting to", p.Name)
			p.UpdateValidMoves(b, moveHistory)
		}
	}
}

func (p *Piece) SelectPromotionPiece() (string, string) {
	// Request player move from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pawn has been promoted - Please choose your promotion piece from (K)night, (B)ishop, (R)ook, (Q)ueen: ")
	inputMove, _ := reader.ReadString('\n')
	inputMove = strings.TrimSpace(inputMove)
	validPromotionPieces := []string{"K", "B", "R", "Q"}
	if !slices.Contains(validPromotionPieces, inputMove) {
		fmt.Println("Invalid promotion piece")
		_, _ = p.SelectPromotionPiece()
	}

	if inputMove == "K" {
		if p.Color == "White" {
			return "Knight", "♘"
		} else {
			return "Knight", "♞"
		}
	} else if inputMove == "B" {
		if p.Color == "White" {
			return "Bishop", "♗"
		} else {
			return "Bishop", "♝"
		}
	} else if inputMove == "R" {

		if p.Color == "White" {
			return "Rook", "♖"
		} else {
			return "Rook", "♜"
		}
	} else if inputMove == "Q" {
		if p.Color == "White" {
			return "Queen", "♕"
		} else {
			return "Queen", "♛"
		}
	}
	return "", ""
}

func (p *Piece) KnightMoves(b *Board) {
	moves := []Position{
		{p.CurrentPosition.Row + 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row + 2, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row - 2, p.CurrentPosition.Col + 1}, {p.CurrentPosition.Row - 2, p.CurrentPosition.Col - 1},
		{p.CurrentPosition.Row + 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row + 1, p.CurrentPosition.Col - 2},
		{p.CurrentPosition.Row - 1, p.CurrentPosition.Col + 2}, {p.CurrentPosition.Row - 1, p.CurrentPosition.Col - 2},
	}

	ValidPositions := make([]Position, 0, len(moves))
	currentColor := p.Color

	for _, pos := range moves {
		if pos.IsInBounds() && pos.CanBeOccupied(b, currentColor) {
			ValidPositions = append(ValidPositions, pos)
		}
	}

	p.ValidMoves = ValidPositions
}

func (p *Piece) RookMoves(b *Board) {

	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)
	rowCells := b.GetRow(currentCell)
	colCells := b.GetCol(currentCell)

	p.ValidMoves = append(
		p.CheckLinearMoves(b, rowCells, currentCell, "row"),
		p.CheckLinearMoves(b, colCells, currentCell, "col")...,
	)
}

func (p *Piece) BishopMoves(b *Board) {
	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)
	p.ValidMoves = p.CheckDiagonalMoves(b, currentCell)
}

func (p *Piece) QueenMoves(b *Board) {
	currentCell := b.GetCell(p.CurrentPosition.Row, p.CurrentPosition.Col)

	// Generate all possible diagonal and linear moves
	diagonals := p.CheckDiagonalMoves(b, currentCell)
	rowPositions := p.CheckLinearMoves(b, b.GetRow(currentCell), currentCell, "row")
	columnPositions := p.CheckLinearMoves(b, b.GetCol(currentCell), currentCell, "col")

	// Combine all valid positions
	p.ValidMoves = append(append(diagonals, rowPositions...), columnPositions...)
}

func (p *Piece) KingMoves(b *Board) {
	directions := []Position{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	p.ValidMoves = []Position{}

	for _, d := range directions {
		newPos := Position{p.CurrentPosition.Row + d.Row, p.CurrentPosition.Col + d.Col}
		if newPos.IsInBounds() && newPos.CanBeOccupied(b, p.Color) && !b.CheckCellAttacked(newPos, p.Color) {
			p.ValidMoves = append(p.ValidMoves, newPos)
		}
	}

	//TODO: Implement Castling (and inability to castle through check)
	// iF King and target rook havent moved (no piece history) and cells between them are not attacked
	// then move position of king then move position of rook
}

func (p *Piece) CheckLinearMoves(board *Board, cells []Cell, currCell Cell, linearType string) []Position {

	outputPosition := []Position{}
	var Curr int
	if linearType == "col" {
		Curr = currCell.Position.Row
	} else if linearType == "row" {
		Curr = currCell.Position.Col
	}

	for i := Curr + 1; i <= 7; i++ {
		if !cells[i].Position.HasPiece(board) {
			outputPosition = append(outputPosition, cells[i].Position)
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color == currCell.Piece.Color {
			break
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color != currCell.Piece.Color {
			outputPosition = append(outputPosition, cells[i].Position)
			break
		}
	}

	for i := Curr - 1; i >= 0; i-- {
		if !cells[i].Position.HasPiece(board) {
			outputPosition = append(outputPosition, cells[i].Position)
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color == currCell.Piece.Color {
			break
		} else if cells[i].Position.HasPiece(board) && cells[i].Piece.Color != currCell.Piece.Color {
			outputPosition = append(outputPosition, cells[i].Position)
			break
		}
	}

	return outputPosition

}

func (p *Piece) CheckDiagonalMoves(board *Board, currCell Cell) []Position {
	diagonals := []Position{}
	currPosition := currCell.Position
	currentRow := currPosition.Row
	currentCol := currPosition.Col

	// get up left
	for i := 1; currentRow+i < 8 && currentCol-i >= 0; i++ {
		if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol - i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol - i})
		} else if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol - i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol - i})
			break
		} else {
			break
		}
	}

	// get up right
	for i := 1; currentRow+i < 8 && currentCol+i < 8; i++ {
		if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol + i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol + i})
		} else if board.GetCellByPosition(Position{Row: currentRow + i, Col: currentCol + i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow + i, Col: currentCol + i})
			break
		} else {
			break
		}
	}
	// get down left
	for i := 1; currentRow-i >= 0 && currentCol-i >= 0; i++ {
		if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol - i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol - i})
		} else if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol - i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol - i})
			break
		} else {
			break
		}
	}
	// get down right
	for i := 1; currentRow-i >= 0 && currentCol+i < 8; i++ {
		if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol + i}).Piece == nil {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol + i})
		} else if board.GetCellByPosition(Position{Row: currentRow - i, Col: currentCol + i}).Piece.Color != currCell.Piece.Color {
			diagonals = append(diagonals, Position{Row: currentRow - i, Col: currentCol + i})
			break
		} else {
			break
		}
	}

	// Sort by row first, then by column
	sort.Slice(diagonals, func(i, j int) bool {
		if diagonals[i].Row != diagonals[j].Row {
			return diagonals[i].Row < diagonals[j].Row
		}
		return diagonals[i].Col < diagonals[j].Col
	})

	return diagonals
}

// Check to see if the opposite color king is in the new list of valid moves for that piece,
// if yes set the game state to 1 (check)
// if the king also cannot move set the game state to 2 (checkmate)
// -- also need to check if any other piece can block the current check
func (p *Piece) UpdateGameState(b *Board, moveHistory []MoveHistory) (bool, bool, bool) {
	color := p.Color

	for _, position := range p.ValidMoves {
		if b.GetCellByPosition(position).Piece != nil && b.GetCellByPosition(position).Piece.Name == "King" && b.GetCellByPosition(position).Piece.Color != color {

			// check if the king can move
			king := b.GetCellByPosition(position).Piece
			king.UpdateValidMoves(b, moveHistory)
			if len(king.ValidMoves) == 0 {
				// check if any other piece can block the check
				// for every piece of the other color
				// if they have a move in valid moves which stops the current pieces' attack then return (true, false, false)
				// if they have no move in valid moves which stops the current pieces' attack then return (false, true, false)
				for _, row := range b.Cells {
					for _, cell := range row {
						if cell.Piece != nil && cell.Piece.Color != color {
							cell.Piece.UpdateValidMoves(b, moveHistory)
							if IsInValidMoves(cell.Piece.ValidMoves, king.ValidMoves) {
								// TODO - need to implement proper logic to only check moves between the attacking piece and the king
								return true, false, false
							}
						}
					}
				}
				// Checkmate as the attacked king cannot move and no piece and block/take the piece
				return false, true, false
			}
			// Check as the attacked king can move
			return true, false, false
		}

		canMove := false
		for _, row := range b.Cells {
			for _, cell := range row {
				if cell.Piece != nil && cell.Piece.Color != color {
					cell.Piece.UpdateValidMoves(b, moveHistory)
					if len(cell.Piece.ValidMoves) > 0 {
						canMove = true
					}
				}
			}
		}
		// Stalemate as the attacked king cannot move and no piece and block/take the piece
		if !canMove {
			return false, false, true
		}
	}

	// check if no pieces of the opposite color can move and king not checked then stalemate

	return false, false, false
}
