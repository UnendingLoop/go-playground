package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

func main() {
 reader := bufio.NewReader(os.Stdin)

 // Reading input for a
 aInput, _ := reader.ReadString('\n')
 aStrings := strings.Fields(aInput)
 a := make([]int, len(aStrings))
 for i, str := range aStrings {
  num, _ := strconv.Atoi(str)
  a[i] = num
 }

 // Reading input for b
 bInput, _ := reader.ReadString('\n')
 bStrings := strings.Fields(bInput)
 b := make([]int, len(bStrings))
 for i, str := range bStrings {
  num, _ := strconv.Atoi(str)
  b[i] = num
 }

 i := 0
 j := 0
 for i < len(a) && j < len(b) {
  if a[i] != b[j] {
   fmt.Println("NO")
   os.Exit(0)
  }
  i++
  j++

  // Finding the start of a new group of items
  for i < len(a) && a[i] == a[i-1] {
   i++
  }
  for j < len(b) && b[j] == b[j-1] {
   j++
  }
 }

 // Checking if the last characters match
 if i != len(a) || j != len(b) {
  fmt.Println("NO")
 } else {
  fmt.Println("YES")
 }
}