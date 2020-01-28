package main

import "fmt"

func main() {
	name, email := "Sanji", "kang.sanjeet1401@gmail.com"
	size := 1.3

	var age int32 = 12
	var isCool = true
	isCool = false
	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
