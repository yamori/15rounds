package pieces

type Bishop struct {
}

func NewBishop() Bishop {
	b := Bishop{}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (b Bishop) CanCapture(ownCol int, ownRow int, tgtCol int, tgtRow int) bool {
	tgtCol_transpose := tgtCol - ownCol
	tgtRow_transpose := tgtRow - ownRow
	return abs(tgtCol_transpose) == abs(tgtRow_transpose)
}

func (b Bishop) ToString() string {
	return "I am a Bishop"
}
