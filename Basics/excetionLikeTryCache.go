package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://reqbin.com/echo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
