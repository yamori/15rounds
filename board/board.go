package board

import (
	"fmt"
	"reflect"
	"strconv"
	"yamori/15rounds/pieces"
)

type board struct {
	field [8][8]pieces.Piece
}

func NewBoard(m [8][8]pieces.Piece) board {
	b := board{
		field: m,
	}
	return b
}

func convertIndicesToCoords(col int, row int) string {
	colMap := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return colMap[col] + strconv.Itoa(row+1)
}

func (b board) PrintBoard() {
	for col, colArr := range b.field {
		for row, v := range colArr {
			if v == nil {
				fmt.Printf("%v:()\n", convertIndicesToCoords(col, row))
			} else {
				fmt.Printf("%v:%v\n", convertIndicesToCoords(col, row), reflect.TypeOf(v))
				fmt.Println(v.ToString())
			}
		}
	}
}

func (b board) ToString() string {
	return "I'm a ChessBoard"
}
