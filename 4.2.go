package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "white"
	colors[1] = "red"
	colors[2] = "black"

	colors[1] = "pink"
	fmt.Println(colors)
}
