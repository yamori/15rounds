package board

type board struct {
}

func NewBoard() board {
	b := board{}
	return b
}

func (b board) ToString() string {
	return "I'm a ChessBoard"
}
