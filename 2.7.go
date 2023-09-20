package main

import "fmt"

func main() {
	var a int = 15
	var b int = 25
	c := true
	v := false
	fmt.Println(!(a > b), "Right, it's", c)
	fmt.Println(!(a <= b), "No, it's", v)
	if a != b {
		fmt.Println(v)
	} else {
		fmt.Println(c)
	}
}
