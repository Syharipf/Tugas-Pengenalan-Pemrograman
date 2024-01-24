package main

import "fmt"

func main() {

	var awal, kelahiran, imigrasi, kematian, emigrasi, akhir int

	fmt.Print("Masukan jumlah penduduk awal, kelahiran, imigrasi, kematian, dan emigrasi : ")
	fmt.Scan(&awal, &kelahiran, &imigrasi, &kematian, &emigrasi)

	akhir = awal + kelahiran + imigrasi - kematian - emigrasi

	fmt.Println("Total penduduk akhir adalah", akhir)
}
