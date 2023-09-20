package main

import "fmt"

func main() {
	var age int = 17
	var adult bool = true
	var underage bool = false
	bonus := 12.5
	malus := 3.14
	allowed := "Access is allowed"
	forbidden := "Access is forbidden"
	age--
	if age > 17 {
		fmt.Println(adult, bonus, allowed)
	} else {
		fmt.Println(underage, malus, forbidden)
	}
	fmt.Println(age)
}
