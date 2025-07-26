package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	left := strings.Builder{}
	right := strings.Builder{}

	for i := 0; i < n; i++ {
		x := strconv.Itoa(i+1) + " "
		if s[i] == 'L' || s[i] == 'l' {
			left.WriteString(x)
		} else {
			var tmp strings.Builder
			tmp.WriteString(x)
			tmp.WriteString(right.String())
			right = tmp
		}
	}

	fmt.Printf("%s0 %s\n", left.String(), right.String())
}
