package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Url = "https://reqres.in/api/users?page=2"

func main() {
	fmt.Println("web request")

	res, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
	var resWriter strings.Builder
	resWriter.Write(content)

	fmt.Println(resWriter.String())

}
