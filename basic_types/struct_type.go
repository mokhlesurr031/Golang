package main

import "fmt"

type Thing struct {
	name string
	age  int
}

func main() {
	//fmt.Println(Thing{})
	//fmt.Println(Thing{"mahin", 23})
	//fmt.Println(Thing{"azhar", 23})

	th := new(Thing)
	th.name = "Bob"
	th.age = 23
	fmt.Println(th)

	th2 := new(Thing)
	th2.name = "Marley"
	th2.age = 23
	fmt.Println(th2)

}
