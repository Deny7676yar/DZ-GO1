package random

import (
	"github.com/Deny7676yar/fibonacci/web"
	"math/rand"
	"time"
)

type RandomStruct struct {
	fibers []web.Fiber
}

func NewRandomStruct(fibers []web.Fiber) *RandomStruct {
	rand.Seed(time.Now().UnixNano())
	return &RandomStruct{fibers}
}

func (rs *RandomStruct) Fib(n int) int {
	fiber := rs.fibers[rand.Intn(len(rs.fibers))]
	return fiber.Fib(n)
}

func (rs *RandomStruct) Name() string {
	fiber := rs.fibers[rand.Intn(len(rs.fibers))]
	return fiber.Name()
}