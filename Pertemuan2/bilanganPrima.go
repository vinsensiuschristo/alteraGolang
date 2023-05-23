package main

import "fmt"

func main() {
	var numberFromUser uint32

	fmt.Print("Masukan angka bilangan yang ingin dicek : ")

	fmt.Scanln(&numberFromUser)

	if numberFromUser > 65535{
		fmt.Println("Angka Tidak Valid. Masukan dari 0 - 65535")
	} else if numberFromUser == 2 {
		fmt.Println("Bilangan yang Anda masukan adalah bilangan prima")
	} else if numberFromUser % 2 == 1 {
		fmt.Println("Bilangan yang Anda masukan adalah bilangan prima")
	} else {
		fmt.Println("Bilangan yang Anda masukan adalah bukan bilangan prima")
	}
}