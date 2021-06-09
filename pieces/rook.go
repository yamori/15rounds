package pieces

type rook struct {
}

func NewRook() rook {
	r := rook{}
	return r
}

func (r rook) ToString() string {
	return "I am a Bishop"
}
