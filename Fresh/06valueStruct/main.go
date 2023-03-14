package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")
	// no inheritance in golang no super no parent

	var anuj = User{"anuj", "anuj@mail.com", true, 22}

	fmt.Println(anuj)
	fmt.Printf("anuj details are: %+v\n", anuj)

	fmt.Printf("anuj name is: %v\n", anuj.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
