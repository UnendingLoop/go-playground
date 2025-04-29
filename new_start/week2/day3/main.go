package main

import (
	"bufio"
	"calculator/calculate"
	"calculator/config"
	"calculator/errorHandler"
	"calculator/input"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	a, b            float64
	operation, expr string
	err             error
)

func init() {
	/*Init используем для приветствия и для проверки наличия
	флагов командной строки, и ипри их наличии - выполнения
	"автовычислений" и передачи управления обратно в main
	для интерактивной работы.
	*/
	flag.StringVar(&expr, "expression", "", "выражение для вычисления")
	flag.Float64Var(&a, "a", 0, "первое число")
	flag.Float64Var(&b, "b", 0, "второе число")
	flag.StringVar(&operation, "operation", "", "операция")
	flag.Parse()
	fmt.Println("Добро пожаловать в Калькулятор!")
	fmt.Println("Программа умеет выполнять операции: сложение, вычитание, умножение, деление и возведение в степень.")

	switch {
	case expr != "":
		fmt.Printf("Введенное выражение: %s\n", expr)
		substrings := strings.Fields(expr)
		if len(substrings) != 3 {
			err = fmt.Errorf("ошибка: %w: %d", config.ErrIncorQuantity, len(substrings))
			errorHandler.ErrorHandler(err)
			return
		}
		a, err = strconv.ParseFloat(substrings[0], 64)
		if err != nil {
			err = fmt.Errorf("ошибка: %w - первое число: %s", config.ErrConvertInput, substrings[0])
			errorHandler.ErrorHandler(err)
			return
		}
		b, err = strconv.ParseFloat(substrings[2], 64)
		if err != nil {
			err = fmt.Errorf("ошибка: %w - второе число: %s", config.ErrConvertInput, substrings[2])
			errorHandler.ErrorHandler(err)
			return
		}
		operation = substrings[1]
		answer, err := calculate.Calculate(a, b, operation)
		if err != nil {
			errorHandler.ErrorHandler(err)
			return
		}
		fmt.Printf("Ответ: %.3f\n", answer)
		return
	case a != 0 || b != 0 || operation != "":
		fmt.Printf("Введенные в терминал данные: %.2f %s %.2f \n", a, operation, b)
		answer, err := calculate.Calculate(a, b, operation)
		if err != nil {
			errorHandler.ErrorHandler(err)
			return
		}
		fmt.Printf("Ответ: %.3f\n", answer)
		return
	default:
		return
	}
}

func main() {
	restart := true
	reader := bufio.NewReader(os.Stdin) // Создаём один раз и передаём в функции
	for restart {
		// Объявляем переменные для хранения двух чисел, операции и ответа
		a, b, operation, err = input.Input(reader) // Просим пользователя ввести операнды
		if err != nil {
			if !errorHandler.ErrorHandler(err) {
				continue
			}
		}
		answer, err := calculate.Calculate(a, b, operation)
		if err != nil {
			if !errorHandler.ErrorHandler(err) {
				continue
			}
		}
		fmt.Printf("Ответ: %f\n", answer)
		fmt.Println("Хотите начать заново? (+/-)")
		want_restart, _ := reader.ReadString('\n')
		want_restart = strings.TrimSpace(want_restart)
		if want_restart == "-" {
			restart = false
			fmt.Println("До свидания!")
		}

	}
}
