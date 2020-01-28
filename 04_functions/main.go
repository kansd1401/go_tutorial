package main

import "fmt"

func greeting(name string) string {
	return "Hellow " + name
}

func main() {
	fmt.Println(greeting("Seymor"))
}
