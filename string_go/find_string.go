package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello how are you hello"
	toFind := "hello"
	fmt.Println(strings.Index(s, toFind))

	fmt.Println(strings.LastIndex(s, toFind))
	fmt.Println(len(s))

	idx := strings.Index(s, toFind)
	fmt.Printf("Index: %d", idx)
}
