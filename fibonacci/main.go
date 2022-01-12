package main

import (

	"fmt"
	"github.com/Deny7676yar/fibonacci/cache"
	"github.com/Deny7676yar/fibonacci/closure"
	"github.com/Deny7676yar/fibonacci/random"
	"github.com/Deny7676yar/fibonacci/recursion"
	"github.com/Deny7676yar/fibonacci/web"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	//fiber := fibImplementations[rand.Intn(len(fibImplementations))]
	//fmt.Printf("Starting %s\n", fiber.Name())
	fiber := random.NewRandomStruct(fibImplementations)
	fmt.Println("Starting Fibonacci (Random)")
	web.Serve(fiber)
}
