package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var iterationsCount int
	fmt.Scan(&iterationsCount)
	rawInput := make([]string, iterationsCount)
	reader := bufio.NewReader(os.Stdin)
	var err error

	for i := range iterationsCount {
		fmt.Printf("Input array %d:\n", i)
		rawInput[i], err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("Couldn't read input!")
			return
		}
	}

	for _, v := range rawInput {
		digits := strings.Fields(v)

		countShips := make(map[string]int, 4)
		for _, v := range digits {
			countShips[v]++
		}
		if countShips["1"] == 4 && countShips["2"] == 3 && countShips["3"] == 2 && countShips["4"] == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
