package main

import (
	"fmt"
	url "net/url"
)

const Url string = "https://lco.dev:3000/learn?cursename=reactjs&paymentid=1324324"

func main() {
	fmt.Println("handling web url")
	result, _ := url.Parse(Url)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	for i, _ := range result.Query() {
		fmt.Println(i)
	}

	PartOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	authUrl := PartOfUrl.String()
	fmt.Println(authUrl)

}
