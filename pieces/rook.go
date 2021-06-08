package pieces

import (
	"fmt"
)

type rook struct {
}

func NewRook() rook {
	b := rook{}
	return b
}

func (e rook) ExamplePrintln() {
	fmt.Printf("Rook's println...\n")
}
