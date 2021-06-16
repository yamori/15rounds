package board

import (
	"fmt"
	"reflect"
	"strconv"
	"yamori/15rounds/pieces"
)

type board struct {
	Field [8][8]pieces.Piece
}

func NewBoard(m [8][8]pieces.Piece) board {
	b := board{
		Field: m,
	}
	return b
}

func (b board) RetrieveRook() (pieces.Rook, int, int) {
	for col, colArr := range b.Field {
		for row, v := range colArr {
			switch v.(type) {
			case pieces.Rook:
				return v.(pieces.Rook), col, row
			}
		}
	}
	return pieces.NewRook(), -1, -1
}

func (b board) RetrieveBishop() (pieces.Bishop, int, int) {
	for col, colArr := range b.Field {
		for row, v := range colArr {
			switch v.(type) {
			case pieces.Bishop:
				return v.(pieces.Bishop), col, row
			}
		}
	}
	return pieces.NewBishop(), -1, -1
}

func ConvertIndicesToCoords(col int, row int) string {
	colMap := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return colMap[col] + strconv.Itoa(row+1)
}

func (b board) PrintBoard() {
	for col, colArr := range b.Field {
		for row, v := range colArr {
			if v == nil {
				fmt.Printf("%v:()\n", ConvertIndicesToCoords(col, row))
			} else {
				fmt.Printf("%v:%v\n", ConvertIndicesToCoords(col, row), reflect.TypeOf(v))
				fmt.Println(v.ToString())
			}
		}
	}
}

func (b board) ToString() string {
	return "I'm a ChessBoard"
}
