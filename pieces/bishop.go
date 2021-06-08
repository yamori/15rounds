package bishop

import (
	"fmt"
)

type bishop struct {
}

func New() bishop {
	b := bishop{}
	return b
}

func (e bishop) LeavesRemaining() {
	fmt.Printf("Bishop's println...\n")
}
