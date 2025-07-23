package calculate

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	testData := []struct {
		a, b       float64
		operation  string
		expected   float64
		shouldFail bool
	}{
		{8, 2, "+", 10, false},
		{10, 6, "-", 4, false},
		{5, 6, "*", 30, false},
		{5, 2, "^", 25, false},
		{30, 6, "/", 5, false},
		{5, 0, "/", 0, true},
		{5, 5, "&", 0, true},
	}
	for _, test := range testData {
		result, err := Calculate(test.a, test.b, test.operation)
		if test.shouldFail {
			if err == nil { //тест должен провалиться, но ошибку от функции не получили
				t.Errorf("Expected error for negative case-test expression %f %s %f, but got none. Smth gone wrong.", test.a, test.operation, test.b)
			}
		} else {
			if err != nil { //тест не должен провалиться, и при этом получили ошибку от функции
				t.Errorf("Unexpected error for positive case-test expression: %f %s %f. Smth gone wrong.", test.a, test.operation, test.b)
			} else if result != test.expected { //результат вычислений не совпал с ожидаемым
				t.Errorf("Calculate (%f %s %f) = %f, expected %f", test.a, test.operation, test.b, result, test.expected)
			}
		}
	}
}
