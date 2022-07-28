package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Convert int to string with strconv.Itoa
	var i1 int = -38
	fmt.Printf("i1: %s\n", strconv.Itoa(i1))

	var i2 int32 = 148
	fmt.Printf("i2: %s %T\n", strconv.Itoa(int(i2)), strconv.Itoa(int(i2)))

	//Convert int to string with fmt.Sprintf
	var i3 int = -38
	s3 := fmt.Sprintf("%d", i3)
	fmt.Printf("i3: %s\n", s3)

	var i4 int32 = 148
	s4 := fmt.Sprintf("%d", i4)
	fmt.Printf("i4: %s\n", s4)

	fmt.Printf("Type of i1, i2, i3, i4: %T %T %T %T", i1, i2, s3, s4)

}
