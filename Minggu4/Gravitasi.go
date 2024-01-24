package main

import "fmt"

func main() {

	var berat, merkurius, venus, bumi, mars, yupiter, saturnus, uranus, neptunus float64

	fmt.Print("Masukan berat tubuh dalam satuan KG : ")
	fmt.Scan(&berat)

	merkurius = 9.8 * 0.38
	venus = 9.8 * 0.91
	bumi = 9.8
	mars = 9.8 * 0.38
	yupiter = 9.8 * 2.37
	saturnus = 9.8 * 0.92
	uranus = 9.8 * 0.89
	neptunus = 9.8 * 1.13

	fmt.Println("Berat tubuh awal =", int(berat), "Kg")
	fmt.Println("Berat tubuh di Merkurius =", int(berat*merkurius), "Kg")
	fmt.Println("Berat tubuh di Venus =", int(berat*venus), "Kg")
	fmt.Println("Berat tubuh di Bumi =", int(berat*bumi), "Kg")
	fmt.Println("Berat tubuh di Mars =", int(berat*mars), "Kg")
	fmt.Println("Berat tubuh di Yupiter =", int(berat*yupiter), "Kg")
	fmt.Println("Berat tubuh di Saturnus =", int(berat*saturnus), "Kg")
	fmt.Println("Berat tubuh di Uranus =", int(berat*uranus), "Kg")
	fmt.Println("Berat tubuh di Neptunus =", int(berat*neptunus), "Kg")
}
