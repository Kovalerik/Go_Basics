package main

import "fmt"

func main() {
	yes := true
	no := false
	contra := yes && no
	maybe := yes || no
	pro := !no
	nope := !yes
	fmt.Println(contra, maybe, pro, nope)
}
