package main

import (
	"fmt"
	"math"
)

//Калькулятор операций с одним операндом и с двумя
func main () {
	var a int
	fmt.Println("Введите число операндов ('1' унарные операции, '2' бинарные операции): ")
	fmt.Scan(&a)

	if a == 1{
		unocalculator()

	}else if a == 2 {
		bicalculator()

	}else {
		fmt.Println("Таких операций нет!")
		main()
	}

}

// Функция бинарный калькулятор
//для операций с двумя числами float64
func bicalculator()  {

	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err_a := fmt.Scanln(&a); err_a != nil {
		fmt.Println("Проверьте тип данных первого числа")
	}


	fmt.Print("Введите второе число: ")
	if _, err_b := fmt.Scanln(&b); err_b != nil {
		fmt.Println("Проверьте тип данных второго числа")
	}

	fmt.Print("Введите арефметическую операцию (+, -, *, /, %, ^): ")
	if _, err_op := fmt.Scanln(&op); err_op != nil {
		fmt.Println("Проверьте арефметическую операции")
	}

	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%":
		result = math.Mod(a, b)
	case "^":
		result = math.Pow(a, b)
	default:
		panic(fmt.Sprint("Операция не определена"))
	}
	fmt.Println(a, op, b, "=", result)
}

//унарный калькулятор для операций с одним числом
func unocalculator () {
	var a float64

	fmt.Print("Введите одно число число: ")
	if _, err_uno := fmt.Scanln(&a); err_uno != nil {
		fmt.Println("Проверьте тип данных второго числа")
	}

	var op string
	fmt.Print("Введите арефметическую операцию (sqrt, log, cbrt): ")
	if _, err_op := fmt.Scanln(&op); err_op != nil {
		fmt.Println("Проверьте арефметическую операции")
	}

	var result float64
	switch op {
	case "sqrt":
		result = math.Sqrt(a)
	case "log":
		result =  math.Log(a)
	case "cbrt":
		result =  math.Cbrt(a)

	}
	fmt.Println(op,".", a, "=", result)
}

