package main

import "fmt"

func main() {

	var p, l, K, L int

	fmt.Print("Masukan nilai p dan l : ")
	fmt.Scan(&p, &l)

	L = p * l
	K = 2*p + 2*l

	fmt.Printf("Keliling dari p=%d dan l=%d adalah %d\n", p, l, K)
	fmt.Printf("Luas dari p=%d dan l=%d adalah %d", p, l, L)
}
