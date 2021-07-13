package main

import "fmt"

func main() {
	var arr = []string{"anand", "anuj", "abhinandan", "akhil", "aman"}

	for _, element := range arr {
		fmt.Println(element)
		for i := 0; i < len(element)-1; i++ {
			fmt.Println(element[i])
		}
		/*if element!="anand" {
			farr=append(farr,element)
		}*/
	}
	//fmt.Println(farr)

}
