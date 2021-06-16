package pieces

type Rook struct {
}

func NewRook() Rook {
	r := Rook{}
	return r
}

func (b Rook) CanCapture(ownCol int, ownRow int, tgtCol int, tgtRow int) bool {
	// TODO, implement this later
	return false
}

func (r Rook) ToString() string {
	return "I am a Rook"
}
