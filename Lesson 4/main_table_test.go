package main

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	table := []struct{
		arg []int
		want []int
	}{
		{[]int{3,2,1,4,5,}, []int{23,2,3,354,}},
		{[]int{1,2,4,6,2,6}, []int{3,35,6,7,8,}},
	}
	for _, entry := range table{
		got := InsertionSort(entry.arg)
		if got != entry.want {
			t.Errorf("for %d got %d want %d", entry.arg, got, entry.want)
		}
	}

}
