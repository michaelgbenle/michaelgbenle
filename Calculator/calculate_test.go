package Calculator

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	testicles := []struct {
		str  []string
		want []float64
	}{

		{[]string{"5*4*3*2", "2/3/4/5", "1+2+3+4", "8-2-4-1"}, []float64{120, 0.03333333333333333, 10, 1}},
		{[]string{"1*2*3*4", "20/4/2/1", "4+3+2+1", "4-3-2-1"}, []float64{24, 2.5, 10, -2}},
		{[]string{"5*2*3*5", "30/2/5/1", "6+6+3+3", "41-21-10-5"}, []float64{150, 3, 18, 5}},
		{[]string{"4*4*4*4", "100/5/10/2", "6+6+3+3", "64-24-10-11"}, []float64{256, 1, 18, 19}},
		{[]string{"6+6+3+3", "64-24-10-11"}, []float64{18, 19}},
		{[]string{"6*6*3*3", "64/2/8/2"}, []float64{324, 2}},
		{[]string{"10+11+12+13", "11-2-4-4", "10*4*2*1", "1000/10/5/5"}, []float64{46, 1, 80, 4}},
		{[]string{"5000/10/10/10"}, []float64{5}},
		{[]string{"1001+2+3+4"}, []float64{1010}},
		{[]string{"1*2*3*4", "4+3+2+1", "4-3-2-1"}, []float64{24, 10, -2}},
	}

	for _, tt := range testicles {
		out := Calculate(tt.str...)
		if reflect.DeepEqual(out, tt.want) != true {
			t.Errorf("Calculate(%v)=%v ;got %v", tt.str, tt.want, out)
		}
		break
	}

}

func TestAddition(t *testing.T) {
	testAdd := []struct {
		add  []float64
		want float64
	}{

		{[]float64{2, 3, 4, 5}, 14},
		{[]float64{5, 5, 4, 5}, 19},
		{[]float64{20, 30, 40, 50}, 140},
	}

	for _, tt := range testAdd {
		out := Addition(tt.add)
		if out != tt.want {
			t.Errorf("Addition(%v)=%v ;got %v", tt.add, tt.want, out)
		}
		break
	}

}

func TestSubtraction(t *testing.T) {
	testSubtract := []struct {
		subtract []float64
		want     float64
	}{

		{[]float64{10, 2, 4, 2}, 2},
		{[]float64{100, 5, 5, 5}, 85},
		{[]float64{20, 10, 4, 5}, 1},
	}

	for _, tt := range testSubtract {
		out := Subtraction(tt.subtract)
		if out != tt.want {
			t.Errorf("Subtraction(%v)=%v ;got %v", tt.subtract, tt.want, out)
		}
		break
	}

}

func TestMultiplication(t *testing.T) {
	testMultiply := []struct {
		multiply []float64
		want     float64
	}{

		{[]float64{10, 2, 4, 2}, 160},
		{[]float64{100, 5, 2, 1}, 1000},
		{[]float64{20, 10, 4, 1}, 800},
	}

	for _, tt := range testMultiply {
		out := Multiplication(tt.multiply)
		if out != tt.want {
			t.Errorf("Multiplication(%v)=%v ;got %v", tt.multiply, tt.want, out)
		}
		break
	}

}

func TestDivision(t *testing.T) {
	testDivide := []struct {
		divide []float64
		want   float64
	}{

		{[]float64{10, 2, 1, 1}, 5},
		{[]float64{100, 5, 2, 5}, 2},
		{[]float64{24, 6, 4, 1}, 1},
	}

	for _, tt := range testDivide {
		out := Division(tt.divide)
		if out != tt.want {
			t.Errorf("Multiplication(%v)=%v ;got %v", tt.divide, tt.want, out)
		}
		break
	}

}
