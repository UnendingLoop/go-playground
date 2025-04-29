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
	a, _ := strconv.Atoi(array[0])
	b, _ := strconv.Atoi(array[1])
	c, _ := strconv.Atoi(array[2])
	d, e, f := float64(a), float64(b), float64(c)
	quadraticEquation(d, e, f)

}

func quadraticEquation(a, b, c float64) {

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
