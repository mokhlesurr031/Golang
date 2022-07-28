package main

import "fmt"

func main() {
	s := "str"
	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Printf("Byte at index %d is '%c' (0x%x)\n", i, c, c)
	}
}
