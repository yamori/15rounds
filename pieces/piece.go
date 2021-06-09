package pieces

import "fmt"

type piece interface {
	ToString() string
}

func PrintName(p piece) {
	fmt.Println(p.ToString())
}
