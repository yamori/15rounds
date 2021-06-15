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

	var fld = [8][8]pieces.Piece{}
	fld[8-1][1-1] = rook   // starts at h1 (8-1)*8 + (1-1)
	fld[3-1][3-1] = bishop // starts at c3 (3-1)*8 + (3-1)
	brd := board.NewBoard(fld)
	fmt.Println(brd.ToString())
	brd.PrintBoard()
}
