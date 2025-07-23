package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	word  string
	count int
}

func cleanText(s string) string {
	s = strings.ToLower(s)
	builder := strings.Builder{}

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsSpace(v) {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}
func topWords(s string, n int) []string {
	pairs := []pair{}
	freqMap := make(map[string]int)

	for v := range strings.FieldsSeq(s) {
		freqMap[v]++
	}
	for k, v := range freqMap {
		pairs = append(pairs, pair{word: k, count: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := []string{}
	for _, v := range pairs {
		result = append(result, v.word)
	}
	if len(result) < n {
		return result
	}
	return result[:n]

}
func main() {
	n := 6
	input := "Hello, world! Go, go, GO..."
	cleanInput := cleanText(input)
	result := topWords(cleanInput, n)
	fmt.Println("The result is:", result)

}
