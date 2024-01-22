package main

import "fmt"

func main() {

	var f, x, y float64

	fmt.Print("Masukan nilai x & y : ")
	fmt.Scan(&x, &y)

	f = 1/(3*x*x+10) + 10*y + 7

	fmt.Printf("f(x,y) dari x=%d & y=%d adalah %f", int(x), int(y), f)
}
