package main

import (
	"fmt"
	"log"
)

func main() {
	// Объявляем переменные для хранения трех чисел и ответа
	var a, b, c, answer float64
	var operation string
	fmt.Println("Введите 3 целых числа через пробел:")

	// Проверяем, удалось ли считать все три числа
	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		fmt.Println("Ошибка: Введите только числа!")
		log.Fatal("Ошибка при вводе чисел:", err) //логируем ошибку и завершаем программу
		return
	}

	fmt.Println("Что с этими числами сделать? Введите один символ '/' , '*' , '-' или '+'")
	fmt.Scan(&operation)
	switch operation {
	case "+":
		answer = a + b + c
		fmt.Println("Ответ:", answer)
	case "-":
		answer = a - b - c
		fmt.Println("Ответ:", answer)
	case "*":
		answer = a * b * c
		fmt.Println("Ответ:", answer)
	case "/":
		if b == 0 || c == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			answer = a / b / c
			fmt.Println("Ответ:", answer)
		}
	default:
		fmt.Println("Неверный символ")
	}
}
