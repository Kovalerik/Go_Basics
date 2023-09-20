package main

import "fmt"

func main() {
	var a float64 = 10.327
	var b int = 56
	var c int = int(a)
	var d float64 = float64(b)
	fmt.Println(c, d)
}
