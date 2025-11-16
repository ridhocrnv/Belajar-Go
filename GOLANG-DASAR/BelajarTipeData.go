package main

import "fmt"

func dataType() {
	fmt.Println("Hello, World!")
	fmt.Println("Usia saya: ", 20)
	fmt.Println("IPS saya semester 1:", 3.55)
	fmt.Println("Apakah saya seorang mahasiswa?", true)
	fmt.Println("Apakah saya suka brokoli?", false)

	// len() menghitung jumlah 'byte' (huruf)
	fmt.Println("Jumlah huruf di 'Alex':", len("Alex"))
	fmt.Println("Jumlah huruf di 'Go':", len("Go"))
}