package main

import (
	"fmt"
	"math/big"
	"os"
)

//2. Задание для продвинутых (необязательное). Написать приложение,
//которое ищет все простые числа от 0 до N включительно.
//Число N должно быть задано из стандартного потока ввода.

func main() {
	var numbInt int64

	fmt.Print("Enter number: ")
	fmt.Fscanln(os.Stdin, &numbInt)

	if big.NewInt(numbInt).ProbablyPrime(0) {
		fmt.Println(numbInt, "- простое число")
	} else {
		fmt.Println(numbInt, "- не простое число")
	}
}
