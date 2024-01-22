package main

import "fmt"

func main() {

	var x, b1, b2 int

	fmt.Print("Masukan dau digit bilangan : ")
	fmt.Scan(&x)

	b1 = x / 10
	b2 = x % 10

	fmt.Printf("Hasil = %d%d%d", b1, x, b2)
}
