package main

import "fmt"

func main() {

	var nilai int
	var BuatTugas, StatusLulus bool

	fmt.Print("Masukan nilai dan apakah mengerjakan tugas? : ")
	fmt.Scan(&nilai, &BuatTugas)

	StatusLulus = (nilai > 55 && BuatTugas) || nilai > 90

	fmt.Printf("Nilai = %d dan mengerjakan tugas? %t, apakah lulus? %t", nilai, BuatTugas, StatusLulus)
}
