package main

import (
	"fmt"
	"log"
)

//0. Объявите структуру с именем Struct и полем s типа string.
//	Для нее (receiver) реализуйте методы: pointerMethod() по указателю,
//	valueMethod() по значению.
//	Методы должны выводить на экран содержимое поля s и имя метода.

type Printer interface {
	valueMethod() string
}

type Struct struct{
	s string
}

func (st Struct) pointerMethod() string{
	fmt.Printf("%s", st.s)
}

func (st Struct) valueMethod() string{
	return fmt.Sprintf("%s", st.s)
}



func main()  {
	pointerStruct := s
}