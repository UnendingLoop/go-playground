package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var line string

func main() {

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		line = input.Text()
	}
	array := strings.Split(line, " ") //создание слайса из всех введенных слов
	A, _ := strconv.Atoi(array[0])
	B, _ := strconv.Atoi(array[1])
	C, _ := strconv.Atoi(array[2])
	a, b, c := float64(A), float64(B), float64(C)
	D := b*b - 4*a*c

	if D > 0 {

		x1 := (-b + math.Sqrt(D)) / (2 * a)

		x2 := (-b - math.Sqrt(D)) / (2 * a)
		if x1 > x2 {
			fmt.Print(x2, " ", x1)
		} else {
			fmt.Print(x1, " ", x2)
		}

	} else if D == 0 {

		x := -b / (2 * a)

		fmt.Print(x)

	}

}
