package main 

import "fmt"

func main() {
  fmt.Println("Selamat Datang didi Di telyu")
  nama := ""
  fmt.Print("Masukkan Nama Anda: ")
  fmt.Scan(&nama)
  fmt.Println("Halo, %s", nama)

}