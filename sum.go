package main

import "fmt"

func main() {
	fmt.Print("Enter the number:")
	var fn int
	fmt.Scan(&fn)
	fmt.Print("Enter the number:")
	var sn int
	fmt.Scan(&sn)
	var sum = fn + sn
	fmt.Print(sum)

}
