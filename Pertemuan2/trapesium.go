package main

import "fmt"

func main() {
	var alas1 float32
	var alas2 float32
	var tinggi float32

	fmt.Print("Masukan panjang Alas Pertama (Alas yang berada di bawah) : ")
	fmt.Scanln(&alas1)

	fmt.Print("Masukan panjang Alas Kedua (Alas yang berada di atas) : ")
	fmt.Scanln(&alas2)

	fmt.Print("Masukan tinggi trapesium : ")
	fmt.Scanln(&tinggi)

	if alas1 == 0 || alas2 == 0 || tinggi == 0 {
		fmt.Println("Inputan Tidak Valid Ulangi Proses, Semua Harus Terisi")
	} else {
		luas := (alas1 + alas2) * tinggi / 2
		fmt.Println("Luas Trapesium adalah : ", luas)
	}
	
}