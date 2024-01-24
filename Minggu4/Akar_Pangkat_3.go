package main

import "fmt"

func main() {

	var x, y int
	var akar bool

	fmt.Print("Masukan nilai x dan y : ")
	fmt.Scan(&x, &y)

	akar = y == (x * x * x)

	fmt.Printf("Apakah %d akar pangkat 3 dari %d? %t", x, y, akar)

}
