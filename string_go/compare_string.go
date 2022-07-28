package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "gone"
	s2 := "Gone"

	//Case-insensitive compare
	if strings.EqualFold(s1, s2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	//Case-sensitive compare
	fmt.Println(s1 == s2)
}
