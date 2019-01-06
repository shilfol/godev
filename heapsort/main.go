package main

import "fmt"

type heap struct {
	L []int
}

func compCP(i, b int, l []int) int {
	c := i*2 + 1

	// check bottom
	if c >= b {
		return i
	}

	// check and search right child
	if c+1 < b && l[c+1] > l[c] {
		c += 1
	}

	// compare child vs parent
	if l[i] < l[c] {
		l[i], l[c] = l[c], l[i]
		return c
	}
	return i
}

func scanHeap(i, b int, l []int) {
	for {
		r := compCP(i, b, l)

		if r == i {
			return
		}
		// continue with child
		i = r
	}
}

func (h *heap) makeHeap() {
	b := len(h.L)
	for i := b; i >= 0; i-- {
		scanHeap(i, b, h.L)
	}
}

func (h *heap) pop() int {
	if len(h.L) <= 0 {
		return -1
	}
	var r int

	r, h.L = h.L[0], h.L[1:]

	if len(h.L) <= 0 {
		return r
	}

	b := len(h.L)
	h.L = append([]int{h.L[b-1]}, h.L[:b-1]...)
	scanHeap(0, b, h.L)

	return r
}

func (h *heap) push(n int) {
	h.L = append(h.L, n)

	i := len(h.L) - 1
	for i >= 0 {
		p := (i - 1) / 2
		if h.L[p] >= h.L[i] {
			break
		}
		h.L[p], h.L[i] = h.L[i], h.L[p]
		i = p
	}
}

func main() {
	h := heap{L: []int{8, 1, 5, 4, 6, 3, 2, 7, 9, 3, 5, 7, 1, 8}}

	// construct Heap
	h.makeHeap()
	fmt.Println(h.L)

	// push
	h.push(10)
	fmt.Println(h.L)
	h.push(15)
	fmt.Println(h.L)
	h.push(20)
	fmt.Println(h.L)

	// pop
	for {
		r := h.pop()
		if r == -1 {
			break
		}
		fmt.Println(r, h.L)
	}
}
