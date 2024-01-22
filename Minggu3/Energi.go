package main

import "fmt"

func main() {
	var e0, c, t, e, e1, cc int

	fmt.Print("Masukan nilai E0, E1, dan c : ")
	fmt.Scan(&e0, &e1, &c)

	t = e0 / e1

	e = e0 - c
	cc = e / e1

	if c == 0 {
		fmt.Println(t)
	} else if c != 0 {
		fmt.Println(cc)
	}
}
