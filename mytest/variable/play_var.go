package variable

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func PlayVar(a int, b int) (int, int) {
	a, b = b, a
	fmt.Println(a, b)
	return a, b
}