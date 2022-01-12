package recursion

import "testing"

func BenchmarkCtx_Fib(b *testing.B) {
	fib := NewRecursion()
	fib.Fib(25)
}
