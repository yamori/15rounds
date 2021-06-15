package pieces

import "fmt"

type Piece interface {
	ToString() string
}

func PrintName(p Piece) {
	fmt.Println(p.ToString())
}
