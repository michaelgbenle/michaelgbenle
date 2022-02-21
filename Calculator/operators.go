package Calculator

func Addition(add []float64) float64 {
	var added float64
	for i := 0; i < len(add); i++ {
		added += add[i]
	}
	return added
}

func Division(divide []float64) float64 {
	divided := divide[0]
	for i := 1; i < len(divide); i++ {
		divided /= divide[i]
	}
	return divided
}

func Multiplication(multiply []float64) float64 {
	multiplied := float64(1)
	for i := 0; i < len(multiply); i++ {
		multiplied *= multiply[i]
	}
	return multiplied
}

func Subtraction(subtract []float64) float64 {
	subtracted := subtract[0]
	for i := 1; i < len(subtract); i++ {
		subtracted -= subtract[i]
	}
	return subtracted
}
