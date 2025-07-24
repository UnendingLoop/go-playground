package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Solution -
func Solution(list []int) string {
	sort.Ints(list)
	groupedList := make(map[int][]int, 0)
	answer := ""
	counter := 0
	groupedList[counter] = append(groupedList[counter], list[0])

	for i := 1; i < len(list); i++ {
		if list[i]-list[i-1] == 1 {
			groupedList[counter] = append(groupedList[counter], list[i])
			continue
		}
		counter++
		groupedList[counter] = append(groupedList[counter], list[i])
	}

	for i := 0; i < len(groupedList); i++ {
		switch {
		case i == 0:
			switch {
			case len(groupedList[i]) == 1:
				answer += fmt.Sprint(groupedList[i][0])
			case len(groupedList[i]) == 2:
				answer += fmt.Sprint(groupedList[i][0], ",", groupedList[i][1])
			case len(groupedList[i]) > 2:
				answer += fmt.Sprint(groupedList[i][0], "-", groupedList[i][len(groupedList[i])-1])
			}
		case len(groupedList[i]) == 1:
			answer += fmt.Sprint(",", groupedList[i][0])
		case len(groupedList[i]) == 2:
			answer += fmt.Sprint(",", groupedList[i][0], ",", groupedList[i][1])
		default:
			answer += fmt.Sprint(",", groupedList[i][0], "-", groupedList[i][len(groupedList[i])-1])
		}
	}
	return answer
}
func anotherSolution(list []int) (s string) {
	l := len(list) - 1 //получаем индекс последнего элемента

	for i, j := 0, 0; i < l; i = j {
		s += strconv.Itoa(list[i])                        //добавляем текущий элемент в ответ
		for j = i; (j < l) && (list[j]+1 == list[j+1]); { //запускаем цикл с условием что текущий и следующий элемент имеют разницу в 1
			j++ //сдвигаем правую границу на 1 столько раз, сколько выполняется условие для след элементов
			//в результате j будет иметь значение последнего индекса группы
		}
		if j-i > 1 { //если в группе более 2х элементов, ставим дефис, а если нет - то запятую
			s += "-"
		} else {
			s += ","
		}
		if i == j { //если внутренний цикл только инициализировался и не выполнились условия, инкрементируем j, а затем и i в следующей итерации внешнего цикла
			j++
		}
	}
	s += strconv.Itoa(list[l])
	return s
}

func main() {
	input := []int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
	fmt.Println(Solution(input))
	fmt.Println(anotherSolution(input))

}
