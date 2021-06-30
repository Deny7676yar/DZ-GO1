package main

import (
	"fmt"
	"os"
)

func fibonachi(n int, mapFib map[int]int) int {

	fibvalue, ok := mapFib[n]
	if ok {
		return mapFib[n]
	}
	if !ok {
		fmt.Printf("%d\tneed to be calculated\n", n)
		fibvalue = fibonachi(n-1, mapFib) + fibonachi(n-2, mapFib)
	}

	mapFib[n] = fibvalue
	return mapFib[n]
}

func main() {
	var mapFib = map[int]int{0: 0, 1: 1, 2: 1}
	var n int

	fmt.Fscanln(os.Stdin, &n)
	fib := fibonachi(n, mapFib)

	fmt.Println(fib)

}
