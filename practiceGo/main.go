package main

import "fmt"

func main() {
	var usr = User{"", "h1@email.com", "", 1, "true"}

	usr.NewMail()
	usr.GetStatus()
	fmt.Println("emil is", usr.Email)
}

type User struct {
	Name   string
	Email  string
	Pass   string
	OneAge int
	Status string
}

func (u User) GetStatus() {

	fmt.Println("it is active", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@mail.com"
	fmt.Println("user mail is", u.Email)

}
