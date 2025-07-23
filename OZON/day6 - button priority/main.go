package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func rangeSumFinder(input []int) int {
	result := 0
	addToSum := false
	for _, v := range input {
		switch v {
		case 1:
			addToSum = true
		case 0:
			addToSum = false
		default:
			if addToSum {
				result += v
			}
		}
	}
	return result
}
func main() {
	n := 0
	fmt.Println("How many checks would you like to process?")
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	answer := []int{}
	for i := range n {
		fmt.Printf("Enter group %d:\n", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
		input = strings.TrimSpace(input)
		temp := strings.FieldsSeq(input)
		numbers := []int{}
		for v := range temp {
			item, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Error converting str to int: %s", err)
			}
			numbers = append(numbers, item)
		}
		answer = append(answer, rangeSumFinder(numbers))
	}

	fmt.Println("Resuts are:")
	for _, v := range answer {
		fmt.Println(v)
	}
}
