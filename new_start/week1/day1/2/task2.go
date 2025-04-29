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

func input(reader *bufio.Reader) (float64, float64) {
	// Считываем ввод пользователя
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка введения данных:", err)
	}
	input = strings.TrimSpace(input)
	substrings := strings.Split(input, " ")
	if len(substrings) != 2 {
		log.Fatal("Неверное количество введенных чисел: ", len(substrings))
	}
	a, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		log.Fatal("Ошибка конвертации строки '", substrings, "' в число: ", err)
	}
	b, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		log.Fatal("Ошибка конвертации строки '", substrings, "' в число: ", err)
	}
	return a, b
}

func calculate(a, b float64, operation string) (float64, bool) {
	var answer float64
	var flag = true
	switch operation {
	case "+":
		answer = a + b
	case "-":
		answer = a - b
	case "*":
		answer = a * b
	case "/":
		if b == 0 {
			log.Fatal("Деление на ноль! Так делать нельзя!")
		}
		answer = a / b
	case "^":
		answer = math.Pow(a, b)
	default:
		fmt.Println("Неверный оператор! Попробуйте ещё раз:")
		flag = false
	}
	return answer, flag
}

func main() {
	var restart bool = true
	reader := bufio.NewReader(os.Stdin) // Создаём один раз и передаём в функции
	for restart {
		// Объявляем переменные для хранения двух чисел и ответа
		var answer float64
		var operation string
		fmt.Println("Введите 2 числа через пробел(используйте точку для дробной части):")

		a, b := input(reader) // Просим пользователя ввести операнды

		fmt.Println("Что с этими числами сделать? Введите один символ '/' , '*' , '-' , '+' или '^' для возведения в степень:")

		flag := false
		for !flag {
			operation, _ = reader.ReadString('\n')
			operation = strings.TrimSpace(operation)
			answer, flag = calculate(a, b, operation)
		}
		fmt.Println("Ответ: ", answer)
		fmt.Println("Хотите начать заново? (+/-)")
		want_restart, _ := reader.ReadString('\n')
		want_restart = strings.TrimSpace(want_restart)
		if want_restart == "-" {
			restart = false
			fmt.Println("До свидания!")
		}

	}
}
