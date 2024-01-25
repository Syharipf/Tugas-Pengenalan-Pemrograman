package main

import "fmt"

func main() {

	var n, i, s, keliling, luas int

	fmt.Print("Masukan jumlah perulangan : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Masukan nilai sisi persegi : ")
		fmt.Scan(&s)

		keliling = 4 * s
		luas = s * s

		fmt.Printf("Sisi persegi aadalah %d, luas dan keliling persegi adalah %d dan %d\n", s, luas, keliling)
	}
}
