package main

import "fmt"

func main() {

	var uang int

	fmt.Print("Masukan nilai uang : ")
	fmt.Scan(&uang)

	fmt.Println(uang/10000, "Lembar")
	uang = uang % 10000

	fmt.Println(uang/5000, "Lembar")
	uang = uang % 5000

	fmt.Println(uang/1000, "Lembar")

}
