package main

import "fmt"

func main() {
	var a float64 = 12.5
	var b float64 = -4.0
	add := b + a
	sub := b - a
	multi := a * a
	div := multi / b
	all := sub + div - add
	fmt.Println(add, sub, multi, div, all)
}
