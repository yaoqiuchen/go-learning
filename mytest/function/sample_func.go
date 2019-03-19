package function

import (
	"fmt"
)

func MultipleReturnTypes(input int) (int, int) {
	sum := 0
	for i := 0; i <= input; i++ {
		sum += i
	}

	return sum, input
}

func NamedReturnTypes(who string, input int) (name string, age int) {
	name = "old " + who
	age = input + 10
	return
}

func DeferFunc(input int) int {
	for i := 0; i < input; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("Before Return")
	return input
}