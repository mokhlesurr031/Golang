package main

import "fmt"

func main() {
	arr := []int{}
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr)
}
