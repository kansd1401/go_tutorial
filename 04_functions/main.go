package main

import "fmt"

func greeting(name string) string {
	return "Hellow " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Seymor"))
	fmt.Println(getSum(1, 2))
}
