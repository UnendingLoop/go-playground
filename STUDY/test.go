package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var line1, line2 string
var array1, array2 []string

func main() {
	input1 := bufio.NewScanner(os.Stdin)
	if input1.Scan() {
		line1 = input1.Text()
	}
	input2 := bufio.NewScanner(os.Stdin)
	if input2.Scan() {
		line2 = input2.Text()
	}
	raw_data1 := strings.Split(line1, " ")
	raw_data2 := strings.Split(line2, " ")

	if len(raw_data1) < len(raw_data2) { //находим массив поменьше
		array1 = raw_data1
		array2 = raw_data2
	} else {
		array1 = raw_data2
		array2 = raw_data1
	}
	if array1[0] != array2[0] || array1[len(array1)-1] != array2[len(array2)-1] { //проверяем равенство первого и последнего элемента
		fmt.Print("NO")
	} else {
		flag := contains(array2, array1)
		if flag {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}

func contains(where, what []string) bool {
	for i := 0; i < len(what); i++ {
		for j := 0; j < len(where); j++ {
			if what[i] == where[j] { // вообще есть indexOf, но непонятно, под какие платформы это пишется, так что по старинке
				break
			}
			if j == len(where)-1 {
				// мы дошли до конца массива, и так и не нашли вхождение - значит, у нас есть элемент, который не входит в where, и нужно вернуть false
				return false
			}
		}
	}
	// ни для одного из элементов не сработал return false, а значит, все они найдены
	return true
}
