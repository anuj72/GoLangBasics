package main

import (
	"fmt"
	"time"
)

func main() {

	var time1 = time.Now()
	fmt.Println(time1)
	fmt.Println(time1.Format("01-02-2006 Monday"))

}
