package main

import (
	"fmt"
	"strings"
)

func main() {
	var arr = []string{"anand", "anuj", "abhinandan", "akhil", "aman"}
	var farr = []string{}
	for _, element := range arr {
		var split = strings.Split(element, "")
		//fmt.Println(split)
		index := 0
		for _, element := range split {
			if element == "a" {
				index++
			}
		}
		if index > 1 {
			farr = append(farr, element)
		}
	}
	fmt.Println(farr)

}
