package main

import "fmt"

var a, b, c, d, x float32

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	if a == 0 || b == 0 {
		fmt.Print("INF")
	} else if a == 0 || b*c == a*d {
		fmt.Print("NO")
	} else if float32(int(b/a))-b/a == 0 {
		x = -b / a
		fmt.Print(x)
	} else {
		fmt.Print("NO")
	}
}
