package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var line1, line2 string
var array1, array2 []string
var flag string

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
		array1 = unique(array1)
		array2 = unique(array2)
		//fmt.Println(array1)
		//fmt.Println(array2)

		if len(array1) == len(array2) {
			for i := 0; i < len(array1); i++ {
				if array1[i] != array2[i] {
					flag = "NO"
					fmt.Print(flag)
					break
				} else {
					flag = "YES"
				}
			}
			fmt.Print(flag)
		} else {
			flag = "NO"
			fmt.Print(flag)
		}
	}
}

func unique(arr []string) []string {
	var new []string
	new = append(new, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			new = append(new, arr[i])
		}
	}
	return new
}
