package main

import "fmt"

func main() {
	fmt.Println("hello")

	d := add1(2, 4)
	print(d)

	e, k := add2(2, 4)
	print(e)
	print(k)

}

//single value return fn
func add1(num1 int, num2 int) int {
	c := num2 + num1
	return c
}

//multy value return fn
func add2(num1 int, num2 int) (int, int) {
	c := num2 + num1
	d := num2 - num1
	return c, d

}
