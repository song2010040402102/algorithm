package main

import (
	"fmt"
)

func sort(a []int) {
	for i := len(a)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] < a[j+1] {
				t := a[j]
				a[j] = a[j+1]
				a[j+1] = t
			}
		}
	}
}

func main() {
	a := []int{1, 2, 5, 3, 9, 4, 5, 6, 1}
	sort(a)
	fmt.Println(a)
}
