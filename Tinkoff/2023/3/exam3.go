package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var line1, line2 string
var counter int

func main() {
	input1 := bufio.NewScanner(os.Stdin)
	if input1.Scan() {
		line1 = input1.Text()
	}
	input2 := bufio.NewScanner(os.Stdin)
	if input2.Scan() {
		line2 = input2.Text()
	}
	arrayA := strings.Split(line1, " ")
	arrayB := strings.Split(line2, " ")
	if len(arrayA) > len(arrayB) {
		counter = len(arrayB)
	} else {
		counter = len(arrayA)
	}
	if arrayA[0] != arrayB[0] || arrayA[len(arrayA)-1] != arrayB[len(arrayB)-1] { //проверяем равенство первого и последнего элемента
		fmt.Print("NO")
	} else {
		a, b := 1, 1
		for i := 0; i <= counter; i++ {
			if a < len(arrayA) && b < len(arrayB) {
				if arrayA[a] != arrayB[b] {
					fmt.Print("NO")
					os.Exit(0)
				}
				a++
				b++
				for { //проскакиваем повторения элементов первого массива
					if a < len(arrayA) && arrayA[a] == arrayA[a-1] {
						a++
					} else {
						break
					}
				}
				for { //проскакиваем повторения элементов второго массива
					if b < len(arrayB) && arrayB[b] == arrayB[b-1] {
						b++
					} else {
						break
					}
				}
			}
		}
		fmt.Print("YES")
	}
}
