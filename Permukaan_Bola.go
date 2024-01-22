package main

import "fmt"

func main() {

	var r, l float32

	fmt.Print("Masukan nilai jari-jari : ")
	fmt.Scan(&r)

	l = 4 * 22 * r * r / 7

	fmt.Printf("Luas permukaan bola dengan r=%f adalah %f", r, l)
}
