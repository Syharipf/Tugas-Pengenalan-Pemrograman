package main

import "fmt"

func main() {

	var i, n, faktorial int

	fmt.Print("Masukan nilai faktorial : ")
	fmt.Scan(&n)

	faktorial = 1

	for i = 1; i <= n; i++ {

		faktorial *= i
	}

	fmt.Printf("%d! adalah %d", n, faktorial)
}
