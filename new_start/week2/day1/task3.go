package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	ErrEmptyInput     = errors.New("пустой ввод")
	ErrIncorrectInput = errors.New("некорректные данные")
	ErrConvertInput   = errors.New("несконвертировался ввод")
	ErrDivideByZero   = errors.New("деление на ноль")
	ErrIncorQuantity  = errors.New("неверное количество введенных чисел")
	ErrIncorrectOper  = errors.New("введена неверная операция")
)

func input(reader *bufio.Reader) (float64, float64, error) {
	fmt.Println("Введите 2 числа через пробел(используйте точку для дробной части):")
	// Считываем ввод пользователя
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %w: %s", ErrIncorrectInput, err)
	}
	// Удаляем символ новой строки и пробелы
	input = strings.TrimSpace(input)
	substrings := strings.Split(input, " ")
	if len(substrings) != 2 {
		return 0, 0, fmt.Errorf("ошибка: %w: %d", ErrIncorQuantity, len(substrings))
	}
	a, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %w - первое число: %s", ErrConvertInput, substrings[0])
	}
	b, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %w - второе число: %s", ErrConvertInput, substrings[1])
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
		return 0, fmt.Errorf("невозможно %w", ErrDivideByZero)
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
		return 0, fmt.Errorf("ошибка: %w", ErrIncorrectOper)
	}
}

func main() {
	var restart bool = true
	reader := bufio.NewReader(os.Stdin) // Создаём один раз и передаём в функции
	for restart {
		// Объявляем переменные для хранения двух чисел, операции и ответа
		var operation string

		a, b, err := input(reader) // Просим пользователя ввести операнды
		if errors.Is(err, ErrIncorrectInput) {
			fmt.Println("❌Ошибка: Введите только числа!") // Выводим сообщение об ошибке
			continue
		} else if errors.Is(err, ErrIncorQuantity) {
			fmt.Println("❌Ошибка: Введите только 2 числа!") // Выводим сообщение об ошибке
			continue
		} else if errors.Is(err, ErrConvertInput) {
			fmt.Println("❌Ошибка: Введите только числа!") // Выводим сообщение об ошибке
			continue
		} else {
			fmt.Printf("✅ Ввод корректен!\n") // Выводим сообщение, что ввод корректен
		}

		flag := false
		for !flag {
			fmt.Println("❓Что с этими числами сделать?\nВведите символ '/' , '*' , '-' , '+' или '^':")
			operation, _ = reader.ReadString('\n')
			operation = strings.TrimSpace(operation)
			answer, err := calculate(a, b, operation)
			if errors.Is(err, ErrDivideByZero) {
				fmt.Println("❌Ошибка: на ноль делить нельзя!")
				continue
			} else if errors.Is(err, ErrIncorrectOper) {
				fmt.Println("❌Ошибка: Введите только один из символов: '/', '*', '-', '+', '^'") // Выводим сообщение об ошибке
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
