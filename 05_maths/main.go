package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	var myNumOne int = 2
	var myNumTwo float64 = 4.4

	fmt.Println("The sum is: ", myNumOne+int(myNumTwo)) // not a valid way

	// random numbers
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
