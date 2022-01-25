package main

import (
    "fmt"
)

type Point struct {
	x, y int
	g, h int
	parent *Point
}

func (p *Point) key() int {
	return p.x << 32 + p.y
}

func (p *Point) val() int {
	return p.g + p.h
}

func (p *Point) next(d uint8) *Point {
	pt := &Point{x: p.x, y: p.y}
	switch d {
	case 1:
		pt.x--
	case 2:
		pt.x++
	case 3:
		pt.y--
	case 4:
		pt.y++
	case 5:
		pt.x--
		pt.y--
	case 6:
		pt.x--
		pt.y++
	case 7:
		pt.x++
		pt.y--
	case 8:
		pt.x++
		pt.y++
	}
	return pt
}

func (p *Point) nextG(d uint8) int {
	switch d {
	case 1, 2, 3, 4:
		return p.g + 10
	case 5, 6, 7, 8:
		return p.g + 14
	}
	return p.g
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 *Point) int {
	if abs(p1.x-p2.x) > abs(p1.y-p2.y) {
		return abs(p1.x-p2.x)*10 + abs(p1.y-p2.y)*4
	} else {
		return abs(p1.x-p2.x)*4 + abs(p1.y-p2.y)*10
	}
}

func pushOpen(pts *[]*Point, mpt map[int]bool, pt *Point) {
	(*pts) = append((*pts), pt)
	cur := len(*pts)-1
	for cur > 0 {
		if (*pts)[cur].val() < (*pts)[cur/2].val() {
			(*pts)[cur], (*pts)[cur/2] = (*pts)[cur/2], (*pts)[cur]
		}
		cur /= 2
	}
	mpt[pt.key()] = true
}

func popOpen(pts *[]*Point, mpt map[int]bool) *Point {
	if len(*pts) == 0 {
		return nil
	}
	ret := (*pts)[0]
	(*pts)[0], (*pts)[len(*pts)-1] = (*pts)[len(*pts)-1], (*pts)[0]
	(*pts) = (*pts)[:len(*pts)-1]
	cur := 0
	for {
		p := -1
		l, r := 2*cur+1, 2*cur+2
		if l < len(*pts) {
			if r < len(*pts) {
				if (*pts)[l].val() < (*pts)[r].val() {
					p = l
				} else {
					p = r
				}
			} else {
				p = l
			}
		} else {
			break
		}
		if (*pts)[p].val() < (*pts)[cur].val() {
			(*pts)[p], (*pts)[cur] = (*pts)[cur], (*pts)[p]
			cur = p
		} else {
			break
		}
	}
	delete(mpt, ret.key())
	return ret
}

func aStar(m [][]uint8, s, e *Point) (ret []*Point) {
	var opens []*Point
	mOpen := make(map[int]bool)
	mClose := make(map[int]bool)
	pushOpen(&opens, mOpen, s)
	for {
		cur := popOpen(&opens, mOpen)
		if cur == nil {
			break
		}
		mClose[cur.key()] = true
		for d := uint8(1); d <= 8; d++ {
			nxt := cur.next(d)
			if nxt.x < 0 || nxt.x >= len(m[0]) || nxt.y < 0 || nxt.y >= len(m) ||
				m[nxt.y][nxt.x] != 0 || mClose[nxt.key()] {
				continue
			}
			if _, ok := mOpen[nxt.key()]; ok {
				if cur.nextG(d) < nxt.g {
					nxt.parent = cur
					nxt.g = cur.nextG(d)
				}
			} else {
				nxt.parent = cur
				nxt.g = cur.nextG(d)
				nxt.h = distance(nxt, e)
				pushOpen(&opens, mOpen, nxt)
				if nxt.key() == e.key() {
					for nxt != nil {
						ret = append(ret, nxt)
						nxt = nxt.parent
					}
					return
				}
			}
		}
	}
	return
}

func main() {
    m := [][]uint8{
    	{0, 0, 0, 0, 0, 0, 0, 0},
    	{0, 0, 0, 0, 1, 1, 0, 0},
    	{0, 0, 0, 0, 1, 0, 0, 0},
    	{0, 0, 0, 0, 0, 0, 0, 0},
    	{0, 0, 0, 0, 1, 0, 0, 0},
    	{0, 0, 0, 0, 0, 0, 0, 0},
    }
    paths := aStar(m, &Point{x:2, y:2}, &Point{x:6, y:2})
    for i := len(paths)-1; i >= 0; i-- {
    	if i > 0 {
    		fmt.Printf("(%d, %d)->", paths[i].x, paths[i].y)
    	} else {
    		fmt.Printf("(%d, %d)\n", paths[i].x, paths[i].y)
    	}
    }
}