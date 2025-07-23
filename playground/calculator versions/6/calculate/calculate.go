package calculate

import (
	"calculator/config"
	"fmt"
	"math"
)

func power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}
func add(a, b float64) (float64, error) {
	return a + b, nil
}
func subtract(a, b float64) (float64, error) {
	return a - b, nil
}
func multiply(a, b float64) (float64, error) {
	return a * b, nil
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("невозможно %w", config.ErrDivideByZero)
	}
	return a / b, nil
}

func Calculate(a, b float64, operation string) (float64, error) {
	operationsMap := map[string]func(float64, float64) (float64, error){
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
		"^": power,
	}
	if opfunc, exist := operationsMap[operation]; exist {
		return opfunc(a, b)
	}
	return 0, fmt.Errorf("ошибка: %w", config.ErrIncorrectOper)
}
