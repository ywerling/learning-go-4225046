package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum * 100) / 100
	fmt.Printf("New sum: %v\n\n", sum)

	fmt.Println("The value of phi is ",math.Phi)

	circleRadius := 15.5
	circomference := circleRadius * 2 * math.Pi
	fmt.Printf("The circomference is %.2f\n",circomference)

}
