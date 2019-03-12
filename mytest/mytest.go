package main

import (
	"fmt"
	"github.com/go-learning/mytest/condition"
)

func main() {
	fmt.Println("Hello, Vincent. hello MyTest")
	//fmt.Println(stringutil.Reverse("!ecnaDetyB nioJ I"))
	//fmt.Println(TryThis())

	condition.TestIf("long long long string")
	condition.TestIf("short ")

	condition.TestFor(3)
	condition.BreakLoop(100)

	condition.TestSwitch(4)
	condition.TestSwitch(10)

	condition.TestSwitch2(1)
	condition.TestSwitch2(2)


}
