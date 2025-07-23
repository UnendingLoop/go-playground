package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertStr2IntArray(input []string) []int {
	output := make([]int, 0, len(input))
	for _, v := range input {
		item, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		output = append(output, item)
	}
	return output
}

// 5 5
// origin: 2 1 3 5 2
// rearrangement: 0 0 + 0 +
// result: 2 1 4 5 3

func rearrangeAppointments(n, m int, input []string) string {
	if m > n || m != len(input) {
		return "x"
	}
	intInput := convertStr2IntArray(input)
	dictionary := make(map[int]int)
	for _, v := range intInput {
		dictionary[v]++
	}
	b := strings.Builder{}
	for _, v := range intInput {
		//узнаём, заняты ли соседние значения в карте
		_, existsNext := dictionary[v+1]
		_, existsPrev := dictionary[v-1]

		switch {
		case !existsPrev && v > 1: //по умолчанию отдаем приоритет на уменьшение значения
			if dictionary[v] > 1 {
				dictionary[v]--
			} else {
				delete(dictionary, v)
			}
			v--
			dictionary[v]++
			b.WriteString(" -")

		case !existsNext && v < n:
			if dictionary[v] > 1 {
				dictionary[v]--
			} else {
				delete(dictionary, v)
			}
			v++
			dictionary[v]++
			b.WriteString(" +")

		default:
			b.WriteString(" 0")
		}
	}

	if len(dictionary) != m { //проверяем, остались ли еще коллизии после перезаписей
		return "x"
	}
	return strings.TrimSpace(b.String()) //удаляем пробел в начале строки
}
func main() {
	//n - кол-во окон
	//m - кол-во пациентов
	//k - кол-во циклов обработки
	var m, n, k int

	fmt.Println("How many checks would you like to process?")
	fmt.Scan(&k)
	reader := bufio.NewReader(os.Stdin)
	answer := []string{}
	for i := range k {
		fmt.Printf("Enter total number of slots and nuber of patients %d:\n", i)
		fmt.Scan(&n, &m)
		fmt.Printf("Enter printed page-group %d:\n", i)
		temp, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
		temp = strings.TrimSpace(temp)
		input := strings.Fields(temp)
		answer = append(answer, rearrangeAppointments(n, m, input))
	}
	fmt.Println("Results are:")
	for _, v := range answer {
		fmt.Println(v)
	}
}
