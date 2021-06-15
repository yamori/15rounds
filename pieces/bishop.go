package pieces

type Bishop struct {
}

func NewBishop() Bishop {
	b := Bishop{}
	return b
}

func (b Bishop) ToString() string {
	return "I am a Bishop"
}
