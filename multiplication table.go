package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the number:")
	var fn int
	fmt.Scan(&fn)
	fmt.Print("multiplication table of ", fn, "\n")
	fmt.Print("---------------------------\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", fn, i, fn*i)
	}
}
