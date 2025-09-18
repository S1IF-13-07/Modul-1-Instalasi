package main

import "fmt"

func main() {
	fmt.Println("Selamat datang di Telyuu")
	nama := ""
	fmt.Print("Masukkan nama kamu:")
	fmt.Scan(&nama)
	fmt.Printf("Hallo %s", nama)
}