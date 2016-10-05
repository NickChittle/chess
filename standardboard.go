package main

import "strconv"

const (
	EMPTY  = iota
	PAWN   = iota
	ROOK   = iota
	KNIGHT = iota
	BISHOP = iota
	QUEEN  = iota
	KING   = iota

	NOCOLOR = 0
	WHITE   = 1
	BLACK   = 2
)

type Piece struct {
	piece int
	color int
}

type StandardBoard struct {
	board [8][8]Piece
}

func MakeStandardBoard() StandardBoard {
	board := StandardBoard{}
	board.board[0] = [8]Piece{
		Piece{ROOK, WHITE}, Piece{KNIGHT, WHITE}, Piece{BISHOP, WHITE},
		Piece{KING, WHITE}, Piece{QUEEN, WHITE},
		Piece{BISHOP, WHITE}, Piece{KNIGHT, WHITE}, Piece{ROOK, WHITE}}
	board.board[1] = [8]Piece{
		Piece{PAWN, WHITE}, Piece{PAWN, WHITE}, Piece{PAWN, WHITE}, Piece{PAWN, WHITE},
		Piece{PAWN, WHITE}, Piece{PAWN, WHITE}, Piece{PAWN, WHITE}, Piece{PAWN, WHITE}}

	board.board[6] = [8]Piece{
		Piece{PAWN, BLACK}, Piece{PAWN, BLACK}, Piece{PAWN, BLACK}, Piece{PAWN, BLACK},
		Piece{PAWN, BLACK}, Piece{PAWN, BLACK}, Piece{PAWN, BLACK}, Piece{PAWN, BLACK}}
	board.board[7] = [8]Piece{
		Piece{ROOK, BLACK}, Piece{KNIGHT, BLACK}, Piece{BISHOP, BLACK},
		Piece{KING, BLACK}, Piece{QUEEN, BLACK},
		Piece{BISHOP, BLACK}, Piece{KNIGHT, BLACK}, Piece{ROOK, BLACK}}
	return board
}

func (b StandardBoard) String() string {
	str := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			str += strconv.Itoa(b.board[i][j].piece) + " "
		}
		str += "\n"
	}
	return str
}

func (b *StandardBoard) GetMoves() []Move {
	moveList := make([]Move, 0)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch b.board[i][j].piece {
			case PAWN:
				newMoves := getPawnMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			case ROOK:
				newMoves := getRookMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			case KNIGHT:
				newMoves := getKnightMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			case BISHOP:
				newMoves := getBishopMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			case QUEEN:
				newMoves := getQueenMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			case KING:
				newMoves := getKingMoves(b, i, j)
				moveList = append(moveList, newMoves...)
			}
		}
	}
	return moveList
}

func getPawnMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	piece := b.board[i][j]

	var dir int
	switch piece.color {
	default:
		panic("WHAT COLOR IS THIS?")
	case WHITE:
		dir = 1
	case BLACK:
		dir = -1
	}
	if b.board[i+dir][j].piece == EMPTY {
		moves = append(moves, Move{Position{i, j}, Position{i + dir, j}})
	}
	if j > 0 && canAttack(b, &b.board[i][j], i+dir, j-1) {
		moves = append(moves, Move{Position{i, j}, Position{i + dir, j - 1}})
	}
	if j < 7 && canAttack(b, &b.board[i][j], i+dir, j+1) {
		moves = append(moves, Move{Position{i, j}, Position{i + dir, j + 1}})
	}
	return moves
}

func getRookMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	return moves
}

func getKnightMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	return moves
}

func getBishopMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	return moves
}

func getQueenMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	return moves
}

func getKingMoves(b *StandardBoard, i int, j int) []Move {
	moves := make([]Move, 0)
	return moves
}

func canAttack(b *StandardBoard, p *Piece, i int, j int) bool {
	if b.board[i][j].color == NOCOLOR {
		return false
	}
	return b.board[i][j].color != p.color
}
