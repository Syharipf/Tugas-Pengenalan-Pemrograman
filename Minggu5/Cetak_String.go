package main

import "fmt"

func main() {

	var i, n int
	var karakter string

	fmt.Print("Masukan jumlah perulangan : ")
	fmt.Scan(&n)

	fmt.Print("Masukan sebuah karakter : ")
	fmt.Scan(&karakter)

	for i = 0; i < n; i++ {
		fmt.Println(karakter)
	}
}
