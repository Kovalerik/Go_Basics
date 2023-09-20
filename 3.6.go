package main

import "fmt"

const one = 1
const eleven = 11.55
const greet = "Доброе утро!"
const lie = false
const tr = true

func main() {
	a := eleven - one
	if a <= 5 {
		fmt.Println(greet, "На улице холодно.", lie)
	} else {
		fmt.Println(greet, "На улице прохладно.", tr)
	}
}
