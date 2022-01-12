package cache

import "testing"

func BenchmarkCtx_Fib(b *testing.B) {
	fib := NewCache()
	fib.Fib(25)
}