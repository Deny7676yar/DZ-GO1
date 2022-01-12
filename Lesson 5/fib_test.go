package main

import (
	"testing"
)

const  N  = 20

var sink int

func BenchmarkFib(b *testing.B)  {
	var mapFib = map[int]int{0: 0, 1: 1, 2: 1}
	var res int
	for i := 0;i < b.N; i++{
		res = fibonachi(N,mapFib)
	}
	sink = res
}

func BenchmarkFibline(b *testing.B){
	var res int
	for i := 0;i < b.N; i++{
		res = fibLine(N)
	}
	sink = res
}
