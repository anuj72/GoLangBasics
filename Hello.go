package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	c := 10
	print(a + b + c)
	array := []string{"sfffs", "ppppp"}
	fmt.Println(array)
	for _, element := range array {
		fmt.Println(element)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	 myadd,mysub := add(100,2)
	 fmt.Println(myadd)
	fmt.Println(mysub)





}

func add(num1 , num2 int) (out1, out2 int) {
         out1 =num1+num2
	     out2 =num1-num2
        return
}
