package main

import "fmt"

func main() {
	var flowers int = 16
	var vases int = 4
	var bags int = 5
	a := flowers + vases
	b := flowers - vases
	c := flowers % bags
	d := flowers / vases
	e := flowers / bags
	f := vases * bags
	j := c + a
	fmt.Println(a, b, c, d, e, f, j)
}
