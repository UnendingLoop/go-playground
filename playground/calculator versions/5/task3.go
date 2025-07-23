package main

import (
	"bufio"
	"errors"
	"flag"
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
	a, b              float64
	operation, expr   string
	err               error
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
			err = fmt.Errorf("ошибка: %w: %d", ErrIncorQuantity, len(substrings))
			errorHandler(err)
			return
		}
		a, err = strconv.ParseFloat(substrings[0], 64)
		if err != nil {
			err = fmt.Errorf("ошибка: %w - первое число: %s", ErrConvertInput, substrings[0])
			errorHandler(err)
			return
		}
		b, err = strconv.ParseFloat(substrings[2], 64)
		if err != nil {
			err = fmt.Errorf("ошибка: %w - второе число: %s", ErrConvertInput, substrings[2])
			errorHandler(err)
			return
		}
		operation = substrings[1]
		answer, err := calculate(a, b, operation)
		if err != nil {
			errorHandler(err)
			return
		}
		fmt.Printf("Ответ: %.3f\n", answer)
		return
	case a != 0 || b != 0 || operation != "":
		fmt.Printf("Введенные в терминал данные: %.2f %s %.2f \n", a, operation, b)
		answer, err := calculate(a, b, operation)
		if err != nil {
			errorHandler(err)
			return
		}
		fmt.Printf("Ответ: %.3f\n", answer)
		return
	default:
		return
	}
}
func input(reader *bufio.Reader) (float64, float64, string, error) {
	fmt.Println("Введите 2 числа через пробел(используйте точку для дробной части):")
	// Считываем ввод пользователя
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w: %s", ErrIncorrectInput, err)
	}
	// Удаляем символ новой строки и пробелы
	input = strings.TrimSpace(input)
	substrings := strings.Fields(input)
	if len(substrings) != 2 {
		return 0, 0, "", fmt.Errorf("ошибка: %w: %d", ErrIncorQuantity, len(substrings))
	}
	a, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w - первое число: %s", ErrConvertInput, substrings[0])
	}
	b, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w - второе число: %s", ErrConvertInput, substrings[1])
	}
	fmt.Println("❓Что с этими числами сделать?\nВведите символ '/' , '*' , '-' , '+' или '^':")
	operation, _ = reader.ReadString('\n')
	operation = strings.TrimSpace(operation)
	return a, b, operation, nil
}

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
		return 0, fmt.Errorf("невозможно %w", ErrDivideByZero)
	}
	return a / b, nil
}

func calculate(a, b float64, operation string) (float64, error) {
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
	return 0, fmt.Errorf("ошибка: %w", ErrIncorrectOper)
}

func errorHandler(err error) bool {
	switch {
	case errors.Is(err, ErrIncorrectInput), errors.Is(err, ErrIncorQuantity), errors.Is(err, ErrIncorrectOper):
		fmt.Println("❌Ошибка: Введены некорректные данные. Повторите попытку.")
	case errors.Is(err, ErrConvertInput):
		fmt.Println("❌Ошибка: Введите только числа!")
	case errors.Is(err, ErrDivideByZero):
		fmt.Println("❌Ошибка: на ноль делить нельзя!")
	default:
		fmt.Printf("✅ Ввод корректен!\n")
		return true
	}
	return false
}
func main() {
	restart := true
	reader := bufio.NewReader(os.Stdin) // Создаём один раз и передаём в функции
	for restart {
		// Объявляем переменные для хранения двух чисел, операции и ответа
		a, b, operation, err = input(reader) // Просим пользователя ввести операнды
		if err != nil {
			if !errorHandler(err) {
				continue
			}
		}
		answer, err := calculate(a, b, operation)
		if err != nil {
			if !errorHandler(err) {
				continue
			}
		}
		fmt.Printf("Ответ: %.2f\n", answer)
		fmt.Println("Хотите начать заново? (+/-)")
		want_restart, _ := reader.ReadString('\n')
		want_restart = strings.TrimSpace(want_restart)
		if want_restart == "-" {
			restart = false
			fmt.Println("До свидания!")
		}

	}
}
