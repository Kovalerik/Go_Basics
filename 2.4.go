package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "42"
	var b string = "7"
	var c, err = strconv.Atoi(a)
	var e, errr = strconv.Atoi(b)
	if err != nil {
	} else {
		fmt.Println(c)
	}
	if errr != nil {
	} else {
		fmt.Println(e)
	}
	s := c + e
	t := c - e
	u := c / e
	fmt.Println(s, t, u)
}
