package main

import (
	"fmt"
	"log"
)

//0. Объявите структуру с именем Struct и полем s типа string.
//	Для нее (receiver) реализуйте методы: pointerMethod() по указателю,
//	valueMethod() по значению.
//	Методы должны выводить на экран содержимое поля s и имя метода.

//1. В main создайте переменную pointerStruct,
//проинициализировав адресом нового экземпляра структуры,
//при этом присвоив полю s значение "pointerStruct".
//	Создайте переменную valueStruct,
//проинициализировав еще одним экземпляром структуры,
//при этом присвоив полю s значение "valueStruct".
//Итак, у вас есть 2 способа обращения и 2 метода,
//итого 4 комбинации. Попробуйте запустить их все и убедитесь, что они работают.
//	Почему получается вызвать pointerMethod для valueStruct и valueMethod для pointerStruct?


//type Printer interface {
//	pointerMethod()
//	valueMethod()
//}

type Struct struct{
	s string
}

func (st &Struct) pointerMethod() string{
	return fmt.Sprintf("%s", st.s)
}

func (st *Struct) valueMethod() string{
	return fmt.Sprintf("%s", st.s)
}

func WriteLog(st Struct) {
	log.Println(st.pointerMethod())
	log.Println(st.valueMethod())
}


func main()  {
	pointerStruct := &Struct{s: "pointerStruct"}
	valueStruct := &Struct{s: "valueStruct"}
	WriteLog(*pointerStruct)
	WriteLog(*valueStruct)

}
