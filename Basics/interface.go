package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	peri() float32
}

func main() {
	mycir := circle{20}
	myrect := rect{20, 20}
	fmt.Println(mycir.area())
	fmt.Println(myrect.area())
	mesure(mycir)
}

func mesure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.peri())
}

func (c circle) area() float32 {
	return math.Pi * c.redius * c.redius
}
func (c rect) area() float32 {
	return c.length * c.width
}
func (c circle) peri() float32 {
	return math.Pi * c.redius * 2
}

type rect struct {
	width  float32
	length float32
}

type circle struct {
	redius float32
}
