package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var answer, line string

func main() {

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		line = input.Text()
	}
	array := strings.Split(line, " ") //создание слайса из всех введенных слов
	for i := 0; i <= len(array)-1; i++ {
		if len(array[i]) > len(answer) {
			answer = array[i]
		}
	}
	fmt.Println(answer)
	fmt.Println(len(answer))

}
