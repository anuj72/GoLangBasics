package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("go network call")
	req, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		panic(err)
	}
	byte, err := io.ReadAll(req.Body)

	var resString strings.Builder

	fmt.Println(string(byte))
	bytecount, _ := resString.Write(byte)
	defer req.Body.Close()

	fmt.Println(bytecount)
	//fmt.Println(resString.String())

	reqBody := strings.NewReader(`
       {
             "courseName":"Golang",
              "price":"0"
       }
      `)
	res, err := http.Post("https://reqres.in/api/users?page=2", "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	byte, _ = io.ReadAll(res.Body)
	fmt.Println("Response is", string(byte))

}
