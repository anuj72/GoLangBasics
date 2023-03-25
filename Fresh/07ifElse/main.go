package main

import "fmt"

func main() {
	fmt.Println("if else")
	num := 10
	if num > 10 {
		fmt.Println("this >10", num)
	} else {
		fmt.Println("this <", num)
	}
}
