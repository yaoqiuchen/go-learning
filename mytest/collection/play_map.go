package condition

import "fmt"

func PlayMap(times int) map[string]int {
	res := map[string]int {
		"Init" : 0,
	}

	for i:=0; i<times; i++ {
		fmt.Println(res["Init"])
	}

	return res
}

func ReturnWhat() (int,int) {
	maps := map[string]int {
		"Vincent" : 1,
		"Delete" : 132,
	}
	delete(maps, "Delete")
	return maps["Vincent"], maps["NoBody"]
}

func ReturnWhat2(name string) int {
	maps := map[string]int {
		"Vincent" : 1,
	}
	res, ok := maps[name]
	fmt.Println("Result ", res, " OK?", ok)
	return res
}