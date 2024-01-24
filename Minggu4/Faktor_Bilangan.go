package main

import "fmt"

func main() {

	var (
		x, y   int
		faktor bool
	)

	fmt.Print("Masukan nilai x dan y : ")
	fmt.Scan(&x, &y)

	faktor = (y % x) == 0

	fmt.Printf("Apakah %d habis membagi %d? %t", x, y, faktor)

}
