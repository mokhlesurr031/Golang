package main

import "fmt"

func main() {
	//declara map
	dict := make(map[string]string)

	dict["name"] = "mahin"
	dict["age"] = "23"

	key := "name"

	fmt.Println(dict)
	fmt.Println(dict[key])

	//map with initialized value

	var mp map[string]string = map[string]string{
		"name": "mahin",
		"age":  "23",
	}
	fmt.Println(mp)
	mp["office"] = "Bista"

	fmt.Println(mp)

	val, ok := mp["nadme"] //if key exists ok return true else false
	fmt.Println(val, ok)
	if ok == true {
		fmt.Println("Paisi")
	} else {
		fmt.Println("Painai")
	}

	fmt.Println("---------------")
	for key, val := range mp {
		fmt.Println(key, val)
	}

	//delete from map
	delete(mp, "office")
	fmt.Println(mp)
}
