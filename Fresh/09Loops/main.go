package main

import "fmt"

func main() {
	fmt.Println("loops in golang")
	days := []string{"s", "m", "t", "w", "t", "f"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		//fmt.Println(days[d])
	}
	for _, j := range days {
		fmt.Println(j)
	}

	var roungeValue = 1
	for roungeValue < 10 {
		if roungeValue == 2 {
			goto fcm
		}
		if roungeValue == 5 {

			break
		}
		fmt.Println(roungeValue)
		roungeValue++
	}

fcm:
	fmt.Println("hello")

}
