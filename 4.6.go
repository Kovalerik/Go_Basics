package main

import "fmt"

func main() {
	var colors [4]string
	colors[0] = "white"
	colors[1] = "red"
	colors[2] = "black"
	colors[3] = "orange"
	fmt.Println(colors[:1], colors[3:])
}
