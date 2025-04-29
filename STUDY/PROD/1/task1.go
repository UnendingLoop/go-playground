package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var answer, line string

func main() {

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		line = input.Text()
	}
	array := strings.Split(line, " ") //создание слайса из всех введенных слов
	for i := 0; i < len(array); i++ {
		flag := IsPalindrome(array[i])
		if flag {
			answer = answer + array[i] + " "
		}

	}
	fmt.Print(answer)

}

func IsPalindrome(str string) bool {

	if len(str) <= 1 {

		return true

	}

	firstRune, _ := utf8.DecodeRuneInString(str)

	lastRune, _ := utf8.DecodeLastRuneInString(str)

	if firstRune != lastRune {

		return false
	}

	if utf8.RuneLen(firstRune) < utf8.RuneCountInString(str)-utf8.RuneLen(lastRune) {
		return IsPalindrome(str[utf8.RuneLen(firstRune) : utf8.RuneCountInString(str)-utf8.RuneLen(lastRune)])
	} else {
		return IsPalindrome(str[utf8.RuneCountInString(str)-utf8.RuneLen(lastRune) : utf8.RuneLen(firstRune)])
	}

}
