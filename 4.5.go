package main

import "fmt"

func main() {
	var numbers [2]int
	numbers[0] = 1
	numbers[1] = 2
	var numbers1 [3]int
	numbers1[0] = 11
	numbers1[1] = 21
	numbers1[2] = 31

	var all [2]int
	all[0] = numbers
	all[1] = numbers1
	fmt.Println(all, numbers, numbers1)
} //не понимаю, как поставить правильно тип
