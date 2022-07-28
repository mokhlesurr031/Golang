package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["mahin"] = 23
	m["mostafa"] = 23

	fmt.Println(m["mahin"])
	if m["mahin"] != 0 {
		fmt.Println("Paisii")
	} else {
		fmt.Println("Paini")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
