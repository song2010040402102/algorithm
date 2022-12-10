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
		a[0], a[i] = a[i], a[0]
		heapify(a[:i], 0)
	}
}

func topN(a []int, n int) []int {
	if n > len(a) {
		n = len(a)
	}
	ret := make([]int, n)
	copy(ret, a[:n])
	if n == len(a) {
		sort(ret)
		return ret
	}
	createMinHeap(ret)
	for i := n; i < len(a); i++ {
		if ret[0] < a[i] {
			ret[0] = a[i]
			heapify(ret, 0)
		}
	}
	for i := len(ret)-1; i > 0; i-- {
		ret[0], ret[i] = ret[i], ret[0]
		heapify(ret[:i], 0)
	}
	return ret
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	fmt.Println(topN(a, 5))
	sort(a)
	fmt.Println(a)
}
