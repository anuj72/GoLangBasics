package main

import "fmt"

func main() {
	fmt.Println("defer in golang")

	defer fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("111111")
}
