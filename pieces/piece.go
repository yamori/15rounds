package pieces

import "fmt"

type Piece interface {
	ToString() string
	CanCapture(ownCol int, ownRow int, tgtCol int, tgtRow int) bool
}

func PrintName(p Piece) {
	fmt.Println(p.ToString())
}
