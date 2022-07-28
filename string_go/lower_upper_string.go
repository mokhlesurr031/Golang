package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Mahin is a good boy"
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToTitle(s))
}
