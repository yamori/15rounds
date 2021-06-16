package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"yamori/15rounds/board"
	"yamori/15rounds/pieces"
)

func flipCoin() bool {
	return rand.Intn(2) == 0
}

func rollTwoDie() int {
	// central limit theorem etc.
	return rand.Intn(6) + rand.Intn(6) + 2
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bishop := pieces.NewBishop()
	pieces.PrintName(bishop)

	rook := pieces.NewRook()
	pieces.PrintName(rook)

	var fld = [8][8]pieces.Piece{}
	fld[8-1][1-1] = rook   // starts at h1 (8-1)*8 + (1-1)
	fld[3-1][3-1] = bishop // starts at c3 (3-1)*8 + (3-1)
	brd := board.NewBoard(fld)

	rounds := 15
	for round := 1; round <= rounds; round++ {
		// retrieve bishop and rook, and their coords
		rk, rk_col, rk_row := brd.RetrieveRook()
		bp, bp_col, bp_row := brd.RetrieveBishop()
		if rk_col < 0 || rk_row < 0 || bp_col < 0 || bp_row < 0 {
			fmt.Fprintf(os.Stderr, "error: %v\n", "pieces not found")
			os.Exit(1)
		}

		// flip coin, move rk
		delta := rollTwoDie()
		if flipCoin() {
			// move up
			fmt.Printf("Coinflip:    up; 2Die: %-2v;", delta)
			brd.Field[rk_col][rk_row] = nil
			brd.Field[rk_col][(rk_row+delta)%8] = rk
		} else {
			// move right
			fmt.Printf("Coinflip: right; 2Die: %-2v;", delta)
			brd.Field[rk_col][rk_row] = nil
			brd.Field[(rk_col+delta)%8][rk_row] = rk
		}

		// check if bishop can capture rook
		bp, bp_col, bp_row = brd.RetrieveBishop()
		if bp_col == -1 && bp_row == -1 {
			fmt.Printf("\n  Rook Captures Bishop, game over.\n")
			break
		}
		rk, rk_col, rk_row = brd.RetrieveRook()
		fmt.Printf(" --> Rook:%v; Bishop:%v;\n", board.ConvertIndicesToCoords(rk_col, rk_row), board.ConvertIndicesToCoords(bp_col, bp_row))
		if bp.CanCapture(bp_col, bp_row, rk_col, rk_row) {
			fmt.Printf("  Bishop Captures Rook, game over.\n")
			break
		}
		if round == rounds {
			fmt.Printf("  Rook has evaded Bishop for %v rounds, game over.\n", rounds)
		}

	}
}
