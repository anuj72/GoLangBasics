package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(" json in go lang")
	EncodeJson()
}

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 299, "Learn.online", "abc1234", []string{"web-dev", "js"}},
		{"golang", 0, "Learn.online", "bcd1234", []string{"web-dev", "go"}},
		{"ReactJs", 299, "Learn.online", "hit1234", nil},
	}

	fjson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", fjson)
	//fmt.Println(string(fjson))

}
