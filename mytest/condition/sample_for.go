package condition

import "fmt"

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