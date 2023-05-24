package main

import "fmt"

func main() {
	isPrime := true
	
	fmt.Println("Program Menentukan Bilangan Prima")

	// Input angka
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	
	if num < 2 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "bukan bilangan prima")
	}
}