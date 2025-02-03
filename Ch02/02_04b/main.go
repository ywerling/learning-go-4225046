package main

import "fmt"

func main() {
	i1, i2, i3 := 13, 35, 87
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3 := 2.3, 5.15, 10.12
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum)	

	total := float64(i1) + f2 
	fmt.Println("Total:", total)	

	product := float64(i2) * f3
	fmt.Println("Product:", product)	

	diff := i3 - i2
	fmt.Println("Difference:", diff)	

	modulo := i3 % i1
	fmt.Println("Modulo:", modulo)	
}
