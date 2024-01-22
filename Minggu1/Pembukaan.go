package main

import "fmt"

func main() {

	var nama, Identitas string
	var NIM int

	fmt.Print("masukan nama, NIM, dan Identitas anda : ")
	fmt.Scan(&nama, &NIM, &Identitas)

	fmt.Println("Nama :", nama)
	fmt.Println("NIM :", NIM)
	fmt.Println("Identitas :", Identitas)
}
