package main

import "fmt"

func main() {
	fmt.Println("Selamat Datang di Telyu")
	nama := ""
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&nama)
	fmt.Printf("Halo, %s", nama)
}
