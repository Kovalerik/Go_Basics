package main

import (
	"fmt"
	"os"
)

func main() {
	var spec string
	var quan int
	fmt.Print("Выберите фрукт:")
	fmt.Fscan(os.Stdin, &spec)
	fmt.Print("Укажите количество:")
	fmt.Fscan(os.Stdin, &quan)

	fmt.Println(spec, quan)
}
