package main

import (
	"fmt"
)

func createMinHeap(a []int) {
	for i := len(a)/2-1; i >= 0; i-- {
		heapify(a, i)
	}
}

func heapify(a []int, i int) {
	l := 2*i+1
	if l >= len(a) {
		return
	}
	r := l+1
	if a[i] < a[l] && (r >= len(a) || a[i] < a[r]) {
		return
	}
	t := a[i]
	if r >= len(a) || a[l] < a[r] {
		a[i] = a[l]
		a[l] = t
		heapify(a, l)
	} else {
		a[i] = a[r]
		a[r] = t
		heapify(a, r)
	}
}

func sort(a []int) {
	if len(a) < 2 {
		return
	}
	createMinHeap(a)
	for i := len(a)-1; i > 0; i-- {
		t := a[0]
		a[0] = a[i]
		a[i] = t
		heapify(a[:i], 0)
	}
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	sort(a)
	fmt.Println(a)
}
