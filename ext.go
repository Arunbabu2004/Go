package main

import "fmt"

func main() {
	var fn int
	var sn int
	fmt.Print("Enter First String: ")
	fmt.Scanln(&fn)
	fmt.Print("Enter Second String: ")
	fmt.Scanln(&sn)
	fmt.Print(fn * sn)
}
