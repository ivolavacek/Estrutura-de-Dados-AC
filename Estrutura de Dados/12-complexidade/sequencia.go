package main

import "fmt"

func main() {
	v := []int{1, 2, 6, 3, 1, 4, 7, 10}
	w := []int{1, 2, 6, 3, 9, 4, 7}

	fmt.Println(sequencia(v))
	fmt.Println(sequencia(w))
}

func sequencia(v []int) int {
	n := len(v)
	if n == 0 {
		return 0
	}

	x := v[0]
	seq, maiorSeq := 1, 1

	for i := 1; i < n; i++ {
		y := v[i]
		if y > x {
			seq++
		}
		if seq > maiorSeq {
			maiorSeq = seq
		}
		if y <= x {
			seq = 1
		}
		x = y
	}

	return maiorSeq
}