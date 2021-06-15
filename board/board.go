package board

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"yamori/15rounds/pieces"
)

type board struct {
	field [64]pieces.Piece
}

func NewBoard(m [64]pieces.Piece) board {
	b := board{
		field: m,
	}
	return b
}

func convertLinearToCoord(n int) string {
	col := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	coldIndx := int(math.Floor(float64(n) / 8.0))
	rmndr := strconv.Itoa((n % 8) + 1)
	return col[coldIndx] + rmndr
}

func (b board) PrintBoard() {
	for i, v := range b.field {
		if v == nil {
			fmt.Printf("%v:()\n", convertLinearToCoord(i))
		} else {
			fmt.Printf("%v:%v\n", convertLinearToCoord(i), reflect.TypeOf(v))
			fmt.Println(v.ToString())
		}
	}
}

func (b board) ToString() string {
	return "I'm a ChessBoard"
}
