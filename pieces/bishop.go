package pieces

type bishop struct {
}

func NewBishop() bishop {
	b := bishop{}
	return b
}

func (b bishop) ToString() string {
	return "I am a Bishop"
}
