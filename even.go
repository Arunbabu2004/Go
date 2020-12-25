package main

import "fmt"

func main() {
	fmt.Print("Enter the limit:")
	var n int
	
	fmt.Scan(&n)
	for i := 1; i <= n; i = i + 2 {
		fmt.Println(i)
	}
}
