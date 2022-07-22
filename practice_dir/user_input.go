package main

import "fmt"

func main() {
	var name string
	fmt.Println("Whats your name")
	fmt.Scanln(&name)
	fmt.Println("Hi " + name)

}
