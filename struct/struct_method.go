package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

//method to get age
func (s Student) getAge() int {
	return s.age
}

//method to calculate total grade
func (g Student) getTotal() int {
	total := 0
	ln := len(g.grades)
	for i := 0; i < ln; i++ {
		total += g.grades[i]
	}
	return total
}

func main() {
	s1 := Student{"Mahin", []int{20, 34, 50}, 25}
	fmt.Println("Struct:", s1)
	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.getAge())
	fmt.Println("Total:", s1.getTotal())
}
