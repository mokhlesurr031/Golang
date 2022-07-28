package main

import "fmt"

func main() {
	//Convert float to string with Sprintf
	var f64 float64 = 1.54
	s := fmt.Sprintf("%f", f64)
	fmt.Printf("f is: '%s' and type is: %T\n", s, s)
}
