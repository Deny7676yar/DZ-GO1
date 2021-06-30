package main

import "fmt"

// InsertionSort - Функция сортировки вставками
func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x

	}
	fmt.Println(arr)
}

func main() {
	slice := []int{21, 2, 33, 11, 3, 45}
	InsertionSort(slice)
}
