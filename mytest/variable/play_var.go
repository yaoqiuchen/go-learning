package variable

import (
	"fmt"
)

func PlayVar(a int, b int) (int, int) {
	a, b = b, a
	fmt.Println(a, b)
	return a, b
}

func PlayBool(input bool)  {
	a := !input
	fmt.Println(a)
}