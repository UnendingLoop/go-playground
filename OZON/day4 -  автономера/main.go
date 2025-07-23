package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func isValidString(input string) (output string) {
	if len(input) < 4 {
		return "-"
	}
	regexp1 := regexp.MustCompile(`^[A-Z][0-9]{2}[A-Z]{2}$`) //A99AA
	regexp2 := regexp.MustCompile(`^[A-Z][0-9][A-Z]{2}$`)    //B9BB

	for i := 0; i < len(input); {
		switch {
		case i+4 <= len(input) && regexp2.MatchString(input[i:i+4]):
			output += (input[i:i+4] + " ")
			i += 4
		case i+5 <= len(input) && regexp1.MatchString(input[i:i+5]):
			output += (input[i:i+5] + " ")
			i += 5
		default:
			return "-"
		}
	}
	return
}

func main() {
	fmt.Println("How many strings you want to check?")
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Scan(&k)
	for i := range k {
		fmt.Printf("String %d:\n", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
		input = strings.TrimSpace(input)
		answer := isValidString(input)
		fmt.Println(answer)
	}
}
