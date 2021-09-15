package main

import (
	"fmt"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func bubblesort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func mergesort(arr []int) []int {
	var less []int
	var equal []int
	var more []int
	if len(arr) > 1 {
		pivot := arr[0]
		for _, x := range arr {
			if x < pivot {
				less = append(less, x)
			}
			if x == pivot {
				equal = append(equal, x)
			}
			if x > pivot {
				more = append(more, x)
			}
		}
		return append(append(mergesort(less), equal...), mergesort(more)...)
	} else {
		return arr
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(mergesort(a))
	b := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(bubblesort(b))
}
