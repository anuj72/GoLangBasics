package main

import "fmt"

func main() {
	fmt.Println("hello")
	book := Book{"kjhbdfchj", "ahgdxuu", 11}
	fmt.Println(book.name)

}

type Book struct {
	name   string
	author string
	page   int
}
