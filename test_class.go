package main

import "fmt"

type Student struct{
	func display(){
		println("hello world")
	}
}

func main(){
	var a = Student
	a.display
}