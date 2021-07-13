package main

func main() {

}

type book interface {
	sportsName() string
}

type Book struct {
	name   string
	author string
	page   int
}
