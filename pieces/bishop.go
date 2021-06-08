package pieces

import (
	"fmt"
)

type bishop struct {
}

func NewBishop() bishop {
	b := bishop{}
	return b
}

func (e bishop) ExamplePrintln() {
	fmt.Printf("Bishop's println...\n")
}
