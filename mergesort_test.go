package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := mergesort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	if !reflect.DeepEqual(want, result) {
		t.Fatalf("want %v, but result is %v", want, result)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergesort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bubblesort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for n := 0; n < b.N; n++ {
		quickSort(list, 0, len(list)-1)
	}
}
