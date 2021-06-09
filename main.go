package main

import (
	"fmt"
	"yamori/15rounds/board"
	"yamori/15rounds/pieces"
)

func main() {
	bishop := pieces.NewBishop()
	pieces.PrintName(bishop)

	rook := pieces.NewRook()
	pieces.PrintName(rook)

	brd := board.NewBoard()
	fmt.Println(brd.ToString())
}
