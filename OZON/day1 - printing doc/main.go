package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validateInputSyntax(n int, input string) error {
	rule1 := regexp.MustCompile(`-\d+-`)
	rule2 := regexp.MustCompile(`[^0-9,-]`)

	switch {
	case len(input) == 0:
		return fmt.Errorf("empty input of printed pages: must be at least one page")
	case n == 0:
		return fmt.Errorf("empty input of total number of pages: must be at least one page")
	case rule1.MatchString(input):
		return fmt.Errorf("'1-2-3' format is forbidden, use only '1-3' with 1 hyphen")
	case rule2.MatchString(input):
		return fmt.Errorf("input contains invalid symbols/characters")
	case strings.Contains(input, ",,"):
		return fmt.Errorf("1 comma can only separate 2 digits")
	case strings.Contains(input, "--"):
		return fmt.Errorf("1 dash can only separate 2 digits")
	case string(input[0]) == "," || string(input[len(input)-1]) == ",":
		return fmt.Errorf("comma cannot start or end the line")
	case strings.Contains(input, ",-") || strings.Contains(input, "-,"):
		return fmt.Errorf("comma and dash can only separate 2 digits")
	}
	return nil
}
func leftToPrint(n int, input string) (answer string) {
	stringCheck := validateInputSyntax(n, input)
	if stringCheck != nil {
		return fmt.Sprintf("Error in input syntax: %s", stringCheck)
	}

	printedPages := make([]bool, n+1) //будем использовать для отслеживания распечатанных страниц: индекс и есть номер страницы
	items := strings.SplitSeq(input, ",")
	for v := range items {
		subItems := strings.Split(v, "-")
		switch len(subItems) {
		case 1:
			k, err := strconv.Atoi(subItems[0])
			if err != nil {
				log.Fatalf("Error converting string to int:%s", err)
			}
			if k > n {
				return fmt.Sprintf("Error: page is #%d, which is outside of total of %d", k, n)
			}
			printedPages[k] = true
		case 2:
			n1, err1 := strconv.Atoi(subItems[0])
			n2, err2 := strconv.Atoi(subItems[1])
			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting string to int:\n%s\n%s", err1, err2)
			}
			if n1 > n2 {
				return fmt.Sprintf("Got %d-%d, but must be %d-%d", n1, n2, n2, n1)
			}
			if n1 > n || n2 > n {
				return "Pagenumber cannot exceed a total number of pages"
			}
			for i := n1; i <= n2; i++ {
				printedPages[i] = true
			}
		}
	}
	answer = stringBuilder(printedPages)
	return
}

func stringBuilder(printedPages []bool) string {
	b := strings.Builder{}
	for start, end := 1, 1; start < len(printedPages); start = end {
		if printedPages[start] { //пропускаем распечатанный номер страницы
			end++
			continue
		}
		for end < len(printedPages) && !printedPages[end] {
			end++
		}
		switch end - start {
		case 0:
			fmt.Fprintf(&b, "%d,", start)
		default:
			fmt.Fprintf(&b, "%d-%d,", start, end-1)
		}
	}
	result := b.String() //удаляем последнюю лишнюю запятую
	return result[:len(result)-1]
}

func main() {
	var n, k int
	fmt.Println("How many checks would you like to process?")
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	answer := []string{}
	for i := range n {
		fmt.Printf("Enter total number of pages for group %d:\n", i)
		fmt.Scan(&k)
		fmt.Printf("Enter printed page-group %d:\n", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ", "") //защита от случайных пробелов
		answer = append(answer, leftToPrint(k, input))
	}

	fmt.Println("Resuts are:")
	for _, v := range answer {
		fmt.Println(v)
	}
}
