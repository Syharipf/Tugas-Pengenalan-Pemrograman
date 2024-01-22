package main

import "fmt"

func main() {

	var celcius, reamur, fahrenheit, kelvin float64

	fmt.Print("Masukan nilai suhu celcius : ")
	fmt.Scan(&celcius)

	reamur = celcius * 4 / 5
	fahrenheit = celcius*9/5 + 32
	kelvin = celcius + 273.15

	fmt.Printf("konversi %f celcius ke reamur adalah %f\n", celcius, reamur)
	fmt.Printf("konversi %f celcius ke fahrenheit adalah %f\n", celcius, fahrenheit)
	fmt.Printf("konversi %f celcius ke kelvin adalah %f", celcius, kelvin)
}
