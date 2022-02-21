package Calculator

import (
	"strconv"
	"strings"
)

var calculated []float64
var add []float64
var subtract []float64
var multiply []float64
var divide []float64

func Calculate(str ...string) []float64 {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "+") {
			addIt := strings.Split(str[i], "+")

			for i := 0; i < len(addIt); i++ {
				a, _ := strconv.Atoi(addIt[i])
				add = append(add, float64(a))

			}
			calculated = append(calculated, Addition(add))

		} else if strings.Contains(str[i], "-") {
			subtractIt := strings.Split(str[i], "-")
			for i := 0; i < len(subtractIt); i++ {
				s, _ := strconv.Atoi(subtractIt[i])
				subtract = append(subtract, float64(s))

			}
			calculated = append(calculated, Subtraction(subtract))

		} else if strings.Contains(str[i], "*") {
			multiplyIt := strings.Split(str[i], "*")
			for i := 0; i < len(multiplyIt); i++ {
				m, _ := strconv.Atoi(multiplyIt[i])
				multiply = append(multiply, float64(m))

			}
			calculated = append(calculated, Multiplication(multiply))
		} else if strings.Contains(str[i], "/") {
			divideIt := strings.Split(str[i], "/")
			for i := 0; i < len(divideIt); i++ {
				d, _ := strconv.Atoi(divideIt[i])
				divide = append(divide, float64(d))
			}
			calculated = append(calculated, Division(divide))
		}

	}
	return calculated
}
