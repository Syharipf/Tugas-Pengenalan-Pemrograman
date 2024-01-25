package main

import "fmt"

func main() {

	var i, n, code, total int
	var rata float64

	fmt.Print("Masukan berapa hari mahasiswa beraltih : ")
	fmt.Scan(&n)

	total = 0

	for i = 0; i < n; i++ {
		fmt.Print("Masukan berapa jam mahasiswa berlatih ngoding : ")
		fmt.Scan(&code)

		total += code
	}

	rata = float64(total) / float64(n)

	fmt.Println("Rata-rata waktu yang dihabiskan untuk berlatih ngoding adalah", rata, "jam")
}
