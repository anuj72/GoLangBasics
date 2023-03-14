package main

import "fmt"

func main() {

	fmt.Println("hello")

	var mPointer *int

	fmt.Println(mPointer)

	var mValue = 12
	var pointer = &mValue

	fmt.Println("pointer is ", pointer)
	fmt.Println("pointer value is ", *pointer)

	*pointer = *pointer + 2

	fmt.Println("pointer value is ", *pointer)

}
