package main

import "fmt"

func main() {

	var t BookT
	t = BookT{"10", "14", 11}
	fmt.Println("Area of tank :", t.Tarea())

}
func (m BookT) Tarea() float64 {

	return float64(2 * m.page)
}

type book interface {
	sportsName() string
}

type BookT struct {
	name   string
	author string
	page   int
}
