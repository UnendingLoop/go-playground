package input

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	// Отключаем стандартный вывод в тестах
	originalStdout := os.Stdout                   // Сохраняем оригинальный stdout
	defer func() { os.Stdout = originalStdout }() // Восстанавливаем после выполнения тестов
	os.Stdout = nil                               //"Отключаем" ст. вывод на время теста

	mockData := []struct {
		expression string
		shouldFail bool
		expectedA  float64
		expectedB  float64
		expectedOp string
	}{
		{"3.5 4.2\n+\n", false, 3.5, 4.2, "+"},
		{"3,5 4.2\n+\n", true, 0, 0, ""},
		{"3.5 4,2\n+\n", true, 0, 0, ""},
		{"3.5 \n+\n", true, 0, 0, ""},
		{"3.5 4.2\n\n", true, 0, 0, ""},
		{"3.5 4.2\n+-\n", true, 0, 0, ""},
		{"3.5 4 5\n+\n", true, 0, 0, ""},
	}
	for _, mockInput := range mockData {
		reader := bufio.NewReader(strings.NewReader(mockInput.expression))
		a, b, c, err := Input(reader)
		if mockInput.shouldFail {
			if err == nil {
				t.Errorf("Expected error, received none. Used data: %s", mockInput.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error, received one:\n%s\n. Used data: %s", err, mockInput.expression)
			} else if a != mockInput.expectedA || b != mockInput.expectedB || c != mockInput.expectedOp {
				t.Errorf("Incorrect data assignment. Input: '%s', expected: '%f %s %f', result: '%f %s %f'",
					mockInput.expression, mockInput.expectedA, mockInput.expectedOp, mockInput.expectedB, a, c, b)

			}

		}
	}

}
