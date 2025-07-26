Дана строка 
s
s из маленьких букв английского алфавита. Найдите количество троек 
(
i
;
j
;
k
)
(i;j;k), что 
i
<
j
<
k
i<j<k и 
s
i
=
<
<
a
>
>
,
s
j
=
<
<
b
>
>
,
s
k
=
<
<
c
>
>
s 
i
​
 =<<a>>,s 
j
​
 =<<b>>,s 
k
​
 =<<c>>.

Входные данные
В единственной строке ввода находится строка 
s
s. 
1
≤
∣
s
∣
≤
1
0
5
1≤∣s∣≤10 
5
 .

Выходные данные
Выведите ответ на задачу.
и я написал уже решение для нее - проверь пожалуйста:
package main

import "fmt"

func main() {
	var b string
	var counter int
	var temp int
	fmt.Scan(&b)
	for i := 0; i < len(b)-1; i++ {
		if b[i]<b[ i+1] {
			temp++
		}
		if temp==3{
			temp=0
			counter++
		}
	}
	fmt.Println(counter)
}