package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func input(reader *bufio.Reader) (float64, float64, error) {
	fmt.Println("Введите 2 числа через пробел(используйте точку для дробной части):")
	// Считываем ввод пользователя
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка введения данных: %v", err)
	}
	input = strings.TrimSpace(input)
	substrings := strings.Split(input, " ")
	if len(substrings) != 2 {
		return 0, 0, fmt.Errorf("неверное количество введенных чисел: %d", len(substrings))
	}
	a, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка конвертации ввода '%s' в число: %v", substrings, err)
	}
	b, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка конвертации ввода '%s' в число: %v", substrings, err)
	}
	return a, b, nil
}

func power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}
func add(a, b float64) (float64, error) {
	return a + b, nil
}
func substract(a, b float64) (float64, error) {
	return a - b, nil
}
func multiply(a, b float64) (float64, error) {
	return a * b, nil
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		log.Println("деление на ноль! так делать нельзя!")
		return 0, fmt.Errorf("деление на ноль! так делать нельзя")
	}
	return a / b, nil
}

func calculate(a, b float64, operation string) (float64, error) {
	operations_map := map[string]func(float64, float64) (float64, error){
		"+": add,
		"-": substract,
		"*": multiply,
		"/": divide,
		"^": power,
	}

	if opfunc, exist := operations_map[operation]; exist {
		return opfunc(a, b)
	} else {
		return 0, fmt.Errorf("неверная операция: %s", operation)
	}
}

func main() {
	var restart bool = true
	reader := bufio.NewReader(os.Stdin) // Создаём один раз и передаём в функции
	for restart {
		// Объявляем переменные для хранения двух чисел, операции и ответа
		var operation string

		a, b, err := input(reader) // Просим пользователя ввести операнды
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
			continue
		}

		flag := false
		for !flag {
			fmt.Println("Что с этими числами сделать? Введите один символ '/' , '*' , '-' , '+' или '^' для возведения в степень:")
			operation, _ = reader.ReadString('\n')
			operation = strings.TrimSpace(operation)
			answer, err := calculate(a, b, operation)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
			fmt.Printf("Ответ: %.3f\n", answer)
			flag = true
		}
		fmt.Println("Хотите начать заново? (+/-)")
		want_restart, _ := reader.ReadString('\n')
		want_restart = strings.TrimSpace(want_restart)
		if want_restart == "-" {
			restart = false
			fmt.Println("До свидания!")
		}

	}
}
