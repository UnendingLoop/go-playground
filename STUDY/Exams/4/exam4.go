package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var line1, line2 string
var flag string
var r, a, answer int

func main() {
	input1 := bufio.NewScanner(os.Stdin)
	if input1.Scan() {
		line1 = input1.Text()
	}
	input2 := bufio.NewScanner(os.Stdin)
	if input2.Scan() {
		line2 = input2.Text()
	}
	input := strings.Split(line1, " ")
	L, _ := strconv.Atoi(input[1])

	raw := strings.Split(line2, " ")
	coordinates := make([]int, len(raw))
	for i, str := range raw {
		num, _ := strconv.Atoi(str)
		coordinates[i] = num
	}

	for i := 0; i < len(coordinates); i++ {
		r = coordinates[i] + L
		answer++
		for j := i; j < len(coordinates); j++ {
			if coordinates[j] <= r {
				i++
			} else {
				break
			}
		}
	}
	fmt.Println(answer)

}
