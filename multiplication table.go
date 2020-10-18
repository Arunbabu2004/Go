package main

import (
	"fmt"
)

func main() {
	var fn int
	fmt.Print("Enter the number:")
	fmt.Scan(&fn)
	fmt.Println("multiplication table of ", fn)
	fmt.Println("---------------------------")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", fn, i, fn*i)
	}
}
