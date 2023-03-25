package main

import "fmt"

const name string = "jhbhb" //public

func main() {
	fmt.Println("hello world")
	var i = 10
	fmt.Println(i)
	j := 11
	fmt.Println(j)
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	var anotherVal int
	fmt.Println(anotherVal)
	fmt.Println("type %T \n", anotherVal)
	const abc = ""
	fmt.Println(name)
	var arr = []int{1, 2, 3, 4, 4, 5}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
	var name string = "1111"
	fmt.Println(name)

	var myNum int32 = 122222
	fmt.Println(myNum)

	fmt.Printf("%T\n", myNum)

}
