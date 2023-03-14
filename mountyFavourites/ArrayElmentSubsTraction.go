package main

import "fmt"

//substract two arrays
func main() {
	arr1 := []string{"anuj", "akhil", "anand", "zeeshan", "shadab"}
	arr2 := []string{"anuj", "zeeshan"}
	farr := subc(arr1, arr2)
	fmt.Println(farr)
}
func subc(arr1 []string, arr2 []string) []string {
	farr := []string{}
	for _, element := range arr1 {
		//fmt.Println(element)
		found := false
		for _, melement := range arr2 {
			if element == melement {
				found = true
			}
		}
		if !found {
			farr = append(farr, element)
		}
	}
	return farr

}
