package main

import "fmt"

func main() {

	var posisi0, posisi1, kecepatan, waktu, rumus int

	fmt.Print("Masukan nilai jarak awal, kecepatan, dan waktu : ")
	fmt.Scan(&posisi0, &kecepatan, &waktu)

	rumus = kecepatan * waktu
	posisi1 = rumus + posisi0

	fmt.Println("Hasil = ", posisi1)
	fmt.Printf("Jarak benda bergerak dengan v=%d, dan t=%d adalah %d, posisi akhir = %d + %d = %d ", kecepatan, waktu, rumus, posisi0, rumus, posisi1)

}
