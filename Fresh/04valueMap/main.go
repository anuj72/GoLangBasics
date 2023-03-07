package main

import "fmt"

func main() {

	fmt.Println("Map in golang")
	var name = make(map[string]string)
	name["A"] = "anuj"
	name["B"] = "Bubbble"
	name["C"] = "Cold"

	fmt.Println("list of all", name)
	fmt.Println("list of all", name["A"])

	delete(name, "A")

	fmt.Println("list of all", name)

	for key, value := range name {
		/*fmt.Println(key)
		fmt.Println(value)*/
		fmt.Printf("key is %v value is %v", key, value)
		println("\n")
	}
}
