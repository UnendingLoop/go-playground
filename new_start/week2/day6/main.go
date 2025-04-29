package main //Сортировка слиянием

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
	for i := range size {
		slice[i] = rng.Intn(max-min+1) + min //присвоение рандомного числа
	}
	return slice
}
func mergeSort(source []int) []int {
	if len(source) <= 1 {
		return source
	}
	mid := len(source) / 2
	left := mergeSort(source[:mid])
	right := mergeSort(source[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
func mergeSortUnRec(source []int) []int {
	n := len(source)
	if n <= 1 {
		return source
	}
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-1; left += 2 * size {
			mid := left + size
			right := min(left+2*size-1, n-1)
			if mid >= right {
				continue // Пропускаем некорректные итерации
			}
			mergeUnRec(source, left, mid, right)
		}
	}
	return source
}
func mergeUnRec(arr []int, left, mid, right int) []int {
	//fmt.Printf("Left %v, mid %v, right %v", left, mid, right)
	leftArr := append([]int{}, arr[left:mid]...)
	rightArr := append([]int{}, arr[mid:right]...)
	i, j, k := 0, 0, left
	for i < mid-left && j < right-mid {
		if leftArr[i] < rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
	}

	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
	return arr
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

	sourceArray := generator(min, max, size) //вызов генератора
	resultArray := sourceArray               //делаем копию исходного, чтобы не изменять его
	resultArray = mergeSort(resultArray)
	fmt.Printf("Исходные параметры массива: \n-мин. зн.: %d\n-макс.зн.: %d\n-размер: %d\n", min, max, size)
	fmt.Printf("Сгенерированный массив: \n%v\n", sourceArray)
	fmt.Printf("Отсортированный рекурсивным слиянием массив: \n%v\n", resultArray)
	fmt.Printf("Отсортированный нерекурсивным слиянием массив: \n%v\n", mergeSortUnRec(sourceArray))
}
