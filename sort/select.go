package main

import (
	"fmt"
)

func sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		k := i
		for j := i+1; j < len(a); j++ {
			if a[j] > a[k] {
				k = j
			}
		}
		t := a[k]
		a[k] = a[i]
		a[i] = t
	}
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	sort(a)
	fmt.Println(a)
}
