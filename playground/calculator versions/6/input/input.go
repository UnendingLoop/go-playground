package input

import (
	"bufio"
	"calculator/config"
	"fmt"
	"strconv"
	"strings"
)

func Input(reader *bufio.Reader) (float64, float64, string, error) {
	fmt.Println("Введите 2 числа через пробел(используйте точку для дробной части):")
	// Считываем ввод пользователя
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w: %s", config.ErrIncorrectInput, err)
	}
	// Удаляем символ новой строки и пробелы
	input = strings.TrimSpace(input)
	substrings := strings.Fields(input)
	if len(substrings) != 2 {
		return 0, 0, "", fmt.Errorf("ошибка: %w: %d", config.ErrIncorQuantity, len(substrings))
	}
	a, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w - первое число: %s", config.ErrConvertInput, substrings[0])
	}
	b, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("ошибка: %w - второе число: %s", config.ErrConvertInput, substrings[1])
	}
	fmt.Println("❓Что с этими числами сделать?\nВведите символ '/' , '*' , '-' , '+' или '^':")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)
	if len(operation) == 0 || len(operation) > 1 {
		return 0, 0, "", fmt.Errorf("ошибка: некорректный ввод символа операции над числами '%s' - ожидался 1 символ", operation)
	}
	return a, b, operation, nil
}
