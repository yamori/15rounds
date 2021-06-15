package pieces

type Rook struct {
}

func NewRook() Rook {
	r := Rook{}
	return r
}

func (r Rook) ToString() string {
	return "I am a Bishop"
}
