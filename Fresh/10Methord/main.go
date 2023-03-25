package main

import "fmt"

func main() {

	var usr = User{"Anuj Sachan", "anuj@mail.com", "Golu@1234"}
	mUser := usr.changeMail()

	fmt.Println(mUser)

}

type User struct {
	Name  string
	Email string
	Pass  string
}

func (u User) changeMail() User {
	u.Email = "a1@mail.com"
	return u
}