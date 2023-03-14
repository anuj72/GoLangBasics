package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []string{"anuj", "arun", "abhi", "akshay"}
	var appendArray = append(arr, "ayush", "shfg")
	var appendArray1 = append(appendArray[1:3])
	fmt.Println(appendArray)
	fmt.Println(appendArray1)
	var arr2 = []int{2, 3, 44, 45, 1}
	sort.Ints(arr2)
	fmt.Println(arr2)
	// same array deletion
	// arr2 = [1 2 3 44 45]
	var index = 2

	arr2 = append(arr2[:index], arr2[index:]...)

}
