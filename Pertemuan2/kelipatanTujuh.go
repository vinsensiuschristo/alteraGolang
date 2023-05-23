package main

import "fmt"

func main() {
	var numberFromUser uint32

	fmt.Print("Masukan angka bilangan yang ingin dicek : ")

	fmt.Scanln(&numberFromUser)

	if numberFromUser > 65535 {
		fmt.Println("Angka Tidak Valid. Masukan dari 0 - 65535")
	} else if numberFromUser % 7 == 0 {
		fmt.Println("Bilangan yang Anda masukan adalah bilangan kelipatan 7")
	} else {
		fmt.Println("Bilangan yang Anda masukan adalah bukan bilangan kelipatan 7")
	}
}