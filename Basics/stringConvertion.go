package main

import (
	"fmt"
	"strconv"
)

func main() {
	var marray = []int{1, 35, 5, 6, 6}

	for i := 0; i < len(marray)-1; i++ {
		fmt.Println(strconv.Itoa(i) + " is index")
		fmt.Println(strconv.FormatInt(int64(i), 10) + " is index")

	}

}
