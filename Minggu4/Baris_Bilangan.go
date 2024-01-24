package main

import "fmt"

func main() {

	var suku1, suku2, suku3, suku4, suku5 int

	fmt.Print("Masukan suku ke-1 dan suku ke-2 : ")
	fmt.Scan(&suku1, &suku2)

	suku3 = suku1 + suku2
	suku4 = suku2 + suku3
	suku5 = suku3 + suku4

	fmt.Printf("Suku ke-1 = %d, suku ke-2 = %d, suku ke-3 = %d, suku ke-4 = %d, suku ke-5 = %d", suku1, suku2, suku3, suku4, suku5)
}
