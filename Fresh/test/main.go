package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("test")

	var T1 = time.Now()
	fmt.Println("time now", T1)
	fmt.Println(T1.Format("02/01/2006 Monday"))

	var arr = []int{1, 2, 3, 4, 5, 6, 6, 7, 7}

	arr = append(arr[:2], arr[3:]...)
	fmt.Println(arr)

	var usr = User{"anuj", "anuj@mail.com", "11111"}
	fmt.Println(usr)

}

type User struct {
	name  string
	email string
	pass  string
}
