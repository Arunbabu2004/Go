package main

import "fmt"

func main() {
	fmt.Print("Enter First String: ")
	var fn int
	fmt.Scanln(&fn)
	fmt.Print("Enter Second String: ")
	var sn int
	fmt.Scanln(&sn)
	var sum = fn + sn
	fmt.Print(sum)
}
