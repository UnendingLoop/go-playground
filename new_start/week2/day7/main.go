package main //Быстрая сортировка

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

func quickSort(array []int, start, end int) {
	if start < end {
		pivot := partition(array, start, end)
		quickSort(array, start, pivot-1)
		quickSort(array, pivot+1, end)
	}
}

func partition(arr []int, low, high int) int {
	// Сортируем три элемента для нахождения медианы
	mid := (low + high) / 2
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	// Пивот — это медиана трех элементов
	pivot := arr[mid]
	arr[mid], arr[high] = arr[high], arr[mid] //Перемещаем пивот в конец для корректной отработки
	i := low - 1

	// Цикл должен идти до high, чтобы переместить элементы
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Возвращаем пивот на его правильную позицию
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
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
	//resultArray := sourceArray               //делаем копию исходного, чтобы не изменять его
	//resultArray = mergeSort(resultArray)
	fmt.Printf("Исходные параметры массива: \n-мин. зн.: %d\n-макс.зн.: %d\n-размер: %d\n", min, max, size)
	fmt.Printf("Сгенерированный массив: \n%v\n", sourceArray)
	quickSort(sourceArray, 0, len(sourceArray)-1)
	fmt.Printf("Быстрая сортировка массива: \n%v\n", sourceArray)

}
