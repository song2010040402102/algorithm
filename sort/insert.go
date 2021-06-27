package main

import (
	"fmt"
)

func sort(a []int) {
	for i := 1; i < len(a); i++ {
		k := a[i]
		j := i-1
		for ; j >= 0; j-- {
			if k <= a[j] {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = k
	}
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	sort(a)
	fmt.Println(a)
}
