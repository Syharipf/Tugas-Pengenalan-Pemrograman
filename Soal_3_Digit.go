package main

import "fmt"

func main() {

	var x, d1, d2, d3 int

	fmt.Print("Masukan tiga digit bilangan : ")
	fmt.Scan(&x)

	d1 = (x / 100) % 10
	d2 = (x / 10) % 10
	d3 = x % 10

	fmt.Printf("x=%d, maka d1=%d, d2=%d, d3=%d", x, d1, d2, d3)
}
