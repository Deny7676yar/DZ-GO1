package closure

import "testing"

func BenchmarkCtx_Fib(b *testing.B) {
	fib := NewClosure()
	fib.Fib(25)
}
