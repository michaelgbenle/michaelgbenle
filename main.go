package main

import (
	"fmt"
	"github.com/michaelgbenle/Calculator/Calculator"
)

func main() {
	fmt.Println(Calculator.Calculate("1*2*3*4", "20/4/2/1", "4+3+2+1", "4-3-2-1"))
	fmt.Println(Calculator.Multiplication([]float64{1, 2, 3, 4}))
	fmt.Println(Calculator.Division([]float64{20, 4, 2, 1}))
	fmt.Println(Calculator.Addition([]float64{4, 3, 2, 1}))
	fmt.Println(Calculator.Subtraction([]float64{4, 3, 2, 1}))

}
