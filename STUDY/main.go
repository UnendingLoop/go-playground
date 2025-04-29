package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var answer int
	var line string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		line = input.Text()
	}
	raw_data := strings.Split(line, " ")
	A, _ := strconv.Atoi(raw_data[0]) //стоимость пакета
	B, _ := strconv.Atoi(raw_data[1]) //МБ в рамках пакета
	C, _ := strconv.Atoi(raw_data[2]) //стоимость 1 МБ вне пакета
	D, _ := strconv.Atoi(raw_data[3]) //план МБ на след. мес.
	if D > B {
		answer = A + (D-B)*C
	} else {
		answer = A
	}
	fmt.Print(answer)
}
