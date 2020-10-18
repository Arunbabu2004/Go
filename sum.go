package main

import "fmt"

func main() {
	var sn int
	var fn int
	fmt.Print("Enter the number:")
	fmt.Scan(&fn)
	fmt.Print("Enter the number:")
	fmt.Scan(&sn)
	fmt.Println(fn + sn)

}
