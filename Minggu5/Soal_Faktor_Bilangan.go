package main

import "fmt"

func main() {

	var i, n int

	fmt.Print("Masukan bilangan bulat positif : ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {

		if n%i == 0 {
			fmt.Printf("%d : True\n", i)
		} else {
			fmt.Printf("%d : False\n", i)
		}
	}
}
