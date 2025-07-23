package school_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	school "school/cmd"
	"school/logger"
	"strings"
	"testing"
)

func TestStudentInfo(t *testing.T) {
	var Logger logger.Log
	var analyzer school.SAnalyzer = school.School{school.MockGrades, &Logger} //импортируем моковые данные

	analyzer.StudentInfo("Daddy")
	fmt.Printf("Tested function output: %v\n", Logger.Output)
	expectedResponse := [4]string{"Information about student: Sam",
		"Their marks: [5 0 4]",
		"Their avg score: 3.00",
		"Most successful in subject: Math"}
	for i := 0; i < len(expectedResponse); i++ {
		if !strings.Contains(Logger.Output[i], expectedResponse[i]) {
			t.Errorf("\nExpected output:\n%q\nActual output:\n%q", expectedResponse[i], Logger.Output[i])
		}
	}
}

func TestBestSubject(t *testing.T) {
	var Logger logger.Log
	var analyzer school.SAnalyzer = school.School{school.MockGrades, &Logger} //сюда импортируем моковые данные
	analyzer.BestSubject()

	expectedResponse := "The most successfully learnt subject is Deutsch, its avg score is 4.20."
	if !strings.Contains(Logger.Output[0], expectedResponse) {
		t.Errorf("\nExpected output:\n%q\nActual output:\n%q", expectedResponse, Logger.Output[0])
	}
}

func TestUnevenStudent(t *testing.T) {
	var analyzer school.SAnalyzer = school.School{school.MockGrades, nil}

	expectedResponse := [2]string{"Student with most uneven marks is Alice, their marks are [0 4 5].",
		"Student with most uneven marks is Sam, their marks are [0 4 5]."}
	response := captureWithReadAll(t, func() { analyzer.UnevenStudent() })
	response = strings.TrimSpace(response)
	responseArray := strings.Split(response, "\n")

	for i := 0; i < len(expectedResponse); i++ {
		if expectedResponse[i] != responseArray[i] {
			t.Errorf("\nExpected output:\n%q\nActual output:\n%q", expectedResponse[i], responseArray[i])
		}
	}

}
func captureWithBuffer(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Unsuccessful attempt to create pipe: %s", err)
	}
	originalSTD := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = originalSTD

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
func captureWithReadAll(t *testing.T, f func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Unsuccessful attempt to create pipe: %v", err)
	}
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout

	input, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("Unsuccessful attempt read r: %v", err)
	}
	return string(input)
}
