package main

import "fmt"

func add(x int, y int) {
	// var result = x+y
	fmt.Println(x + y)
}

func main() {
	add(2, 3)
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
}
