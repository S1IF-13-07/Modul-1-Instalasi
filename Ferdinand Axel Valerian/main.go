package main

import "fmt"

func main() {
    fmt.Println("Selamat Datang di Tel-U")
    var nama string
    fmt.Print("Nama saya adalah: ")
    fmt.Scan(&nama)
    fmt.Printf("Halo, %s\n", nama)
}
