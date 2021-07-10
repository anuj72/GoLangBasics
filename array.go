package main

import "fmt"

func main() {
	fmt.Println("hello")
	var arr = []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for index, element := range arr {
		fmt.Println(index, element)
	}

	for _, element := range arr {
		fmt.Println(element)
	}

}
