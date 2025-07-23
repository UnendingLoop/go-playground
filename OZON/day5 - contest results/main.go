package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func numberFinder(input string) (answer string) {
	rule := regexp.MustCompile(`[A-Za-z,](\d+)[A-Za-z,]`)
	arrarr := rule.FindAllStringSubmatch(input, -1)
	for _, v := range arrarr {
		answer += (v[1] + " ")
	}
	return answer
}

func main() {
	n := 0
	fmt.Println("How many strings would you like to check?")
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	for i := range n {
		fmt.Printf("Enter the %dst row:\n", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
			return
		}
		fmt.Println(numberFinder(input))
	}
}
