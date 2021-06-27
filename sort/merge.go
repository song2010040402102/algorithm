package main

import (
	"fmt"
)

func merge(a, b, m []int) {
	if len(a) > 1 {
		merge(a[:len(a)/2], a[len(a)/2:], m[:len(a)])
		copy(a, m[:len(a)])
	}
	if len(b) > 1 {
		merge(b[:len(b)/2], b[len(b)/2:], m[len(a):])
		copy(b, m[len(a):])
	}
	i, j := 0, 0
	for {
		if i >= len(a) && j >= len(b) {
			break
		}
		if j >= len(b) || i < len(a) && a[i] > b[j] {
			m[i+j] = a[i]
			i++
		} else {
			m[i+j] = b[j]
			j++
		}
	}
}

func sort(a []int) {
	if len(a) <= 1 {
		return
	}
	s := make([]int, len(a))
	copy(s, a)
	merge(s[:len(s)/2], s[len(s)/2:], a)
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	sort(a)
	fmt.Println(a)
}
