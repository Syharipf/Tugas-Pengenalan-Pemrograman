package main

import "fmt"

func main() {

	var barang1, barang2, barang3, harga1, harga2, harga3 float64

	fmt.Print("Masukan tiga harga beli barang : ")
	fmt.Scan(&barang1, &barang2, &barang3)

	harga1 = (barang1 * 5 / 100) + barang1
	harga2 = (barang2 * 5 / 100) + barang2
	harga3 = (barang3 * 5 / 100) + barang3

	fmt.Printf("Harga beli = %d, harga jual =%f\n", int(barang1), harga1)
	fmt.Printf("Harga beli = %d, harga jual =%f\n", int(barang2), harga2)
	fmt.Printf("Harga beli = %d, harga jual =%f\n", int(barang3), harga3)
}
