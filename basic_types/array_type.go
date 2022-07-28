package main

import "fmt"

func main() {
	var arr = []int{2, 3}
	fmt.Println(arr)
	fmt.Println(arr[0])

	arr = append(arr, 34)
	fmt.Println(arr)

}
