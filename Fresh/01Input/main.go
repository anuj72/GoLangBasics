package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var welcome = "Welcome pls input"

	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter value :")

	//comma ok //error ok syntexhe
	input, _ := reader.ReadString('\n')
	fmt.Println("input value : " + input)
	var covVal, _ = strconv.ParseFloat(input, 64)
	fmt.Printf("input type : %T", covVal)
	fmt.Println("\n")
	fmt.Printf("input type : %T", input)

	fmt.Println("hello world11")

	fmt.Println("////")

}
