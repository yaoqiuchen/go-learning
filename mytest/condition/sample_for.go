package condition

import (
	"fmt"
)

func TestFor(times int) {
	sum := 0
	for i := 1; i <= times; i++ {
		sum += i
	}
	fmt.Println("sum ", sum)
}

func BreakLoop(times int)  {
	sum := 0
	loop: for i := 1; i <= times; i++ {
		sum += i
		if sum > 40 {
			fmt.Println("BreakLoop: sum bigger than 40")
			break loop
		}
	}
	fmt.Println("BreakLoop: ", sum)
}

func LoopString(input string) {
	for i := 0; i < len(input); i++ {
		fmt.Printf("char at %d is %c, ascII is %d\n", i, input[i], input[i])
	}
	for i, s := range input {
		fmt.Printf("character at %d is %c\n", i, s)
	}
}