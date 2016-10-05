package main

type Position struct {
	x int
	y int
}

type Move struct {
	start Position
	end   Position
}

type Board interface {
	GetMoves() []Move
	MakeMove(Move)
	UndoMove()
}
