package main

import "fmt"

func main() {

	var n, m, i int

	fmt.Print("Masukan 2 bilangan : ")
	fmt.Scan(&n, &m)

	for i = n; i <= m; i++ {
		fmt.Print(i)
	}
}
