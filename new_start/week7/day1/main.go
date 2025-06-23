package main

import (
	"fmt"
	"reflect"
)

type person struct {
	ID   int64
	Name string
	Age  int8
}

func identifier(input interface{}) {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch t.Kind() {
	case reflect.Struct:
		fmt.Printf("It is a structure: '%v' with fields:\n", t.Name())
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("Field %d: '%v', its value '%v'\n", i, t.Field(i).Name, v.Field(i))
			if t.Field(i).Type.Kind() == reflect.Int8 {

			}
		}
	case reflect.Int:
		fmt.Printf("It is an integer with value '%v'\n", v)
	case reflect.String:
		fmt.Printf("It is a string with value '%s'\n", v)
	}

}
func main() {
	human := person{
		ID:   1234567890,
		Name: "Vasiliy",
		Age:  88,
	}

	identifier(human)
	identifier(56)
	identifier("ASdfgh123")

}
