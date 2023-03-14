package main

import "fmt"

func main() {
	fmt.Println("hello")

	var arr = []int{1, 2, 3, 4, 5}

	for index, data := range arr {
		fmt.Println(index, data)
	}
	var myname = fulllname(name{fname: "anuj", lname: "sachan"})
	fmt.Println(myname)
}

type name struct {
	fname string
	lname string
}

func fulllname(name2 name) string {
	return name2.fname + name2.lname
}
