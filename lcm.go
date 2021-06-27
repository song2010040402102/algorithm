package main

import "fmt"

func getCommonM(a int, b int) int {
    if a <= 0 || b <= 0 {
        return 0
    }
    return a*b/getCommonC(a, b)
}

func getCommonC(a int, b int) int {
    ca := getC(a)
    cb := getC(b)
    fmt.Println(ca, cb)
    ret, ia, ib := 1, 0, 0
    for ia < len(ca) && ib < len(cb) {
        if ca[ia] > cb[ib] {
            ib++
        } else if ca[ia] < cb[ib] {
            ia++
        } else {
            ret *= ca[ia]
            ia++
            ib++
        }
    }
    return ret
}

func getC(n int) []int {
    ret := []int{}
    i := 2;
    for i <= n {
        if n%i == 0 {
            n /= i
            ret = append(ret, i)
        } else {
            i++
        }
    }
    return ret
}

func main() {
    fmt.Println(getCommonM(9, 81))
}