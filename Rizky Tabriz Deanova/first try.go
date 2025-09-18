package main

import "fmt"

func main() {
	fmt.Println("Hello, welcome to Telyup!")
	var name string = ""
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s", name)
}
