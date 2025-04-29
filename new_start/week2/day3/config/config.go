package config

import "errors"

var (
	ErrEmptyInput     = errors.New("пустой ввод")
	ErrIncorrectInput = errors.New("некорректные данные")
	ErrConvertInput   = errors.New("несконвертировался ввод")
	ErrDivideByZero   = errors.New("деление на ноль")
	ErrIncorQuantity  = errors.New("неверное количество введенных чисел")
	ErrIncorrectOper  = errors.New("введена неверная операция")
)
