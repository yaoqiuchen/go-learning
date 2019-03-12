package condition

import (
	"fmt"
)

func TestSwitch(times int) {
	switch  {
	case times < 5:
		fmt.Println("input less than 5")
	case times >= 7:
		fmt.Println("input bigger than 5")
	}
}

func TestSwitch2(times int) {
	switch times {
	case 1:
		fmt.Println("input is 1")
	case 2:
		fmt.Println("input is 2")
	}
}

func DetectType() {

}