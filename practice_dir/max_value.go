package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var x int = 2
	var y int = 3
	var result int
	result = max(x, y)
	fmt.Println(result)
}
