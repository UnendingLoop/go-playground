package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var line0, line1, line2 string
var answer, coupon int

func main() {
	input0 := bufio.NewScanner(os.Stdin)
	if input0.Scan() {
		line0 = input0.Text()
		//coupon, _ = strconv.Atoi(line0[0])
	}

	input1 := bufio.NewScanner(os.Stdin)
	if input1.Scan() {
		line1 = input1.Text()
	}

	input2 := bufio.NewScanner(os.Stdin)
	if input2.Scan() {
		line2 = input2.Text()
	}
	raw0 := strings.Split(line0, " ")
	raw1 := strings.Split(line1, " ")
	raw2 := strings.Split(line2, " ")

	coupon := make([]int, len(raw0))
	for i, str := range raw0 {
		num, _ := strconv.Atoi(str)
		coupon[i] = num
	}

	group1 := make([]int, len(raw1))
	for i, str := range raw1 {
		num, _ := strconv.Atoi(str)
		group1[i] = num
	}
	group2 := make([]int, len(raw2))
	for i, str := range raw2 {
		num, _ := strconv.Atoi(str)
		group2[i] = num
	}

	for i := 0; i < len(group1); i++ {
		for j := 0; j < len(group2); j++ {
			if group1[i]+group2[j] == coupon[0] {
				answer++
			}
		}
	}
	fmt.Print(answer)

}
