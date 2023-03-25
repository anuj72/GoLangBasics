package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var time1 = time.Now()
	fmt.Println(time1)
	fmt.Println(time1.Format("01-02-2006 Monday"))

	var v1 = make(map[string]string)
	v1["swjjn"] = "whnchue"
	v1["jhevh"] = "whnchue"

	for index, value := range v1 {
		fmt.Println(index)
		fmt.Println(value)
	}

	var i = "10.55"
	var covVal, _ = strconv.ParseFloat(i, 64)
	fmt.Printf("input type : %T", covVal)

	fmt.Println(i)

	var str = "abc"
	var m1 = pointer(&str)

	fmt.Println(m1)

}

func pointer(t1 *string) string {
	var result = ""
	for _, j := range *t1 {
		result = string(j) + result
	}
	return result

}


