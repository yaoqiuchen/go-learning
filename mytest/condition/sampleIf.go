package condition

import "fmt"

func TestIf(s string) {
	if len(s) >= 10 {
		fmt.Println("string is too long")
	} else {
		fmt.Println("short string")
	}

}