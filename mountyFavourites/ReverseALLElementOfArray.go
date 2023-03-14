package main

import (
	"fmt"
	"strings"
)

func main() {
	var arr = []string{"anand", "anuj", "abhinandan", "akhil", "aman"}
	var farr = []string{}
	for _, element := range arr {
		fmt.Println(element)
		splitted_arr := strings.Split(element, "")
		str := ""
		for i := len(splitted_arr) - 1; i >= 0; i-- {
			str = str + splitted_arr[i]
		}
		farr = append(farr, str)
	}
	fmt.Println(farr)

}
