package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countUnique(input string) int {
	uniqDictionary := make(map[string]struct{}) //пустая структура занимает меньше памяти чем bool
	croppedInput := ""
	for i, j := 0, 0; i < len(input); i = j {
		for j < len(input) && input[i] == input[j] {
			j++
		}
		switch {
		case j-i > 1:
			croppedInput += input[i : i+2]
		default:
			croppedInput += input[i : i+1]
		}
	}
	strings := strings.FieldsSeq(croppedInput)
	for v := range strings {
		uniqDictionary[v] = struct{}{}
	}

	return len(uniqDictionary)
}

func main() {
	n := 0
	fmt.Println("How many checks would you like to process?")
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	answer := []int{}
	for i := range n {
		fmt.Printf("Enter string-group %d:\n", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
		input = strings.TrimSpace(input)
		answer = append(answer, countUnique(input))
	}

	fmt.Println("Resuts are:")
	for _, v := range answer {
		fmt.Println(v)
	}
}
