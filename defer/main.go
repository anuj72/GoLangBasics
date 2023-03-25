package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//defter

func main() {

	fmt.Println("its a file")

	content := "this need a file - learn.com"
	// create os file
	file, err := os.Create("./defer/myLearn.txt")

	if err != nil {
		panic(err)
	}
	// writer for os
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	readFile("./defer/myLearn.txt")

}

func readFile(fileName string) {
	value, error := ioutil.ReadFile(fileName)
	checkNilError(error)
	fmt.Println("text data", string(value))
}

func checkNilError(error error) {
	if error != nil {
		panic(error)
	}
}
