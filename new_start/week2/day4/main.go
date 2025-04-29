package main //практика вызова метода

import (
	"fmt"
	"math"
)

type Data struct {
	a, b      float64
	operation string
}

func (d *Data) Operation() string {
	switch d.operation {
	case "+":
		return fmt.Sprintf("Сумма двух чисел равна %.2f", d.a+d.b)
	case "-":
		return fmt.Sprintf("Если из первого вычесть второе, то будет %.2f", d.a-d.b)
	case "*":
		return fmt.Sprintf("Умножение этих чисел даст %.2f", d.a*d.b)
	case "/", ":", "%":
		if d.b == 0 {
			return "Деление на ноль запрещено!"
		} else { //switch вложен в switch - очень элегантное избежание if-ов
			switch d.operation {
			case "/":
				return fmt.Sprintf("Целая часть от деления равна %.0f", math.Floor(d.a/d.b))
			case ":":
				return fmt.Sprintf("Обычное деление равно %.2f", d.a/d.b)
			case "%":
				return fmt.Sprintf("Дробная часть от деления равна %.4f", d.a/d.b-math.Floor(d.a/d.b))
			}
		}
	case "^":
		return fmt.Sprintf("Результат возведения в степень: %.2f", math.Pow(d.a, d.b))
	default:
		return fmt.Sprintf("Введенная операция '%s' не поддерживаецца!", d.operation)
	}
	return "Unexpected return. Check code."
}

func main() {
	operationsArray := []string{
		"+", "-", "*", ":", "/", "%", "^", ")",
	}
	calc := Data{a: 9, b: 3}
	for i := 0; i < len(operationsArray); i++ {
		calc.operation = operationsArray[i]
		fmt.Println(calc.Operation())
	}
}
