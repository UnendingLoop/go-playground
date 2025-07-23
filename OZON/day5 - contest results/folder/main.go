package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type runner struct {
	id    int
	time  int
	place int
}

func assignPlaces(runnersArr []runner) (answer string) {
	//сортируем по времени финиша - по возрастанию, без лишней логики для одинаковых результатов
	sort.Slice(runnersArr, func(i, j int) bool {
		return runnersArr[i].time < runnersArr[j].time
	})
	runnersArr[0].place = 1 //сразу устанавливаем 1е место первому бегуну в массиве

	for i := 1; i < len(runnersArr); i++ { //начинаем со 2го бегуна
		if runnersArr[i].time-runnersArr[i-1].time < 2 {
			runnersArr[i].place = runnersArr[i-1].place
		} else {
			runnersArr[i].place = i + 1
		}
	}

	sort.Slice(runnersArr, func(i, j int) bool { //сортируем обратно по порядку из входного массива
		return runnersArr[i].id < runnersArr[j].id
	})
	for i := range runnersArr {
		answer += fmt.Sprint(runnersArr[i].place, " ") //формируем ответную строку
	}
	return answer

}

func main() {
	n := 0
	answer := []string{}
	fmt.Println("How many cycles would you like to run?")
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	for range n {
		runnersList := []runner{}
		k := 0
		fmt.Printf("Enter the number of runners:\n")
		fmt.Scan(&k)
		fmt.Printf("Enter their time results:\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
			return
		}
		input = strings.TrimSpace(input)
		for i, v := range strings.Fields(input) {
			time, _ := strconv.Atoi(v)
			runnersList = append(runnersList, runner{id: i, time: time})
		}
		answer = append(answer, assignPlaces(runnersList))
	}
	for _, v := range answer {
		fmt.Println(v)
	}
}
