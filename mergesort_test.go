package main

import (
	"reflect"
	"sort"
	"testing"
)

func makeDescendingIntRange(n int) []int {
	arr := make([]int, n)
	for i := n; i > 0; i-- {
		arr[n-i] = i
	}
	return arr
}

func makeAscendingIntRange(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func TestMergeSort(t *testing.T) {
	to_sort := makeDescendingIntRange(10)
	want := makeAscendingIntRange(10)
	result := mergesort(to_sort)
	if !reflect.DeepEqual(want, result) {
		t.Fatalf("want %v, but result is %v", want, result)
	}
}

func benchmarkGoSort(i int, b *testing.B) {
	to_sort := makeDescendingIntRange(i)
	for n := 0; n < b.N; n++ {
		sort.Ints(to_sort)
	}
}

func BenchmarkGoSort10(b *testing.B) {
	benchmarkGoSort(10, b)
}
func BenchmarkGoSort100(b *testing.B) {
	benchmarkGoSort(100, b)
}

func benchmarkMergeSort(i int, b *testing.B) {
	to_sort := makeDescendingIntRange(i)
	for n := 0; n < b.N; n++ {
		mergesort(to_sort)
	}
}

func BenchmarkMergeSort10(b *testing.B) {
	benchmarkMergeSort(10, b)
}

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSort(100, b)
}

func benchmarkBubbleSort(i int, b *testing.B) {
	to_sort := makeDescendingIntRange(i)
	for n := 0; n < b.N; n++ {
		bubblesort(to_sort)
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	benchmarkBubbleSort(10, b)
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSort(100, b)
}
func benchmarkQuickSort(i int, b *testing.B) {
	to_sort := makeDescendingIntRange(i)
	list := to_sort
	for n := 0; n < b.N; n++ {
		quickSort(list, 0, len(list)-1)
	}
}

func BenchmarkQuickSort10(b *testing.B) {
	benchmarkQuickSort(10, b)
}
func BenchmarkQuickSort100(b *testing.B) {
	benchmarkQuickSort(100, b)
}
