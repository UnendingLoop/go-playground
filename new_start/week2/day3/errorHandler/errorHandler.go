package errorHandler

import (
	"calculator/config"
	"errors"
	"fmt"
)

func ErrorHandler(err error) bool {
	switch {
	case errors.Is(err, config.ErrIncorrectInput), errors.Is(err, config.ErrIncorQuantity), errors.Is(err, config.ErrIncorrectOper):
		fmt.Println("❌Ошибка: Введены некорректные данные. Повторите попытку.")
	case errors.Is(err, config.ErrConvertInput):
		fmt.Println("❌Ошибка: Введите только числа!")
	case errors.Is(err, config.ErrDivideByZero):
		fmt.Println("❌Ошибка: на ноль делить нельзя!")
	default:
		fmt.Printf("✅ Ввод корректен!\n")
		return true
	}
	return false
}
