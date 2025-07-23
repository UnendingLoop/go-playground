package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generator(min, max, size int) []int {
	//создание генератора
	source := rand.NewSource(time.Now().UnixNano()) //задает источник рандомного значения
	rng := rand.New(source)

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rng.Intn(max-min+1) + min //присвоение рандомного числа
	}
	return slice
}

func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		swap := false
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return array
}

func extremumFinder(array []int) (int, int) {
	min, max := array[0], array[0]

	for _, op := range array {
		if op > max {
			max = op
		}
		if op < min {
			min = op
		}
	}
	return min, max
}

func uniquesAndDuplicates(array []int) ([]int, []int) { //фукция поиска дубликатов и уникальных элементов в массиве
	uniqueArray := []int{}
	duplicateArray := []int{}
	counts := make(map[int](int))
	for _, value := range array {
		counts[value]++
	}
	for key, count := range counts {
		if count > 1 {
			duplicateArray = append(duplicateArray, key)
		} else {
			uniqueArray = append(uniqueArray, key)
		}
	}
	return uniqueArray, duplicateArray
}

func main() {
	//Ввод данных и их обработка
	fmt.Println("Введите минимум, максимум, и кол-во элементов(больше 0) для генерации массива через пробел:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода! Перезапустите программу и введите данные снова!")
		return
	}
	input = strings.TrimSpace(input)
	substrings := strings.Fields(input)
	if len(substrings) != 3 {
		fmt.Println("Введено некорректное кол-во переменных! Попробуйте снова!")
		return
	}

	min, err1 := strconv.Atoi(substrings[0])
	max, err2 := strconv.Atoi(substrings[1])
	size, err3 := strconv.Atoi(substrings[2])
	//Общая ошибка
	if err1 != nil || err2 != nil || err3 != nil || size <= 0 {
		fmt.Println("Ошибка ввода! Перезапустите программу и введите данные повторно!")
		return
	}
	switch {
	case max == min:
		fmt.Println("Минимум и максимум совпадают - невозможно сгенерировать массив. Попробуйте снова.")
		return
	case max < min:
		fmt.Println("Минимум и максимум содержат противоречивые данные. Поменяю их местами.")
		min, max = max, min
	}

	sourceArray := generator(min, max, size)    //вызов генератора
	minEl, maxEl := extremumFinder(sourceArray) //вызов поиска макс и мин значения
	uniqueArray, duplicateArray := uniquesAndDuplicates(sourceArray)
	fmt.Printf("Исходные параметры массива: \n-мин. зн.: %d\n-макс.зн.: %d\n-размер: %d\n", min, max, size)
	fmt.Printf("Сгенерированный массив: \n%v\n", sourceArray)
	fmt.Printf("Минимальное значение в массиве: %d, максимальное: %d\n", minEl, maxEl)
	fmt.Printf("Сортировка пузырьком: \n%v\n", bubbleSort(sourceArray))
	fmt.Printf("Массив из уникальных элементов: \n%v\nМассив из повторов: \n%v\n", uniqueArray, duplicateArray)

}
