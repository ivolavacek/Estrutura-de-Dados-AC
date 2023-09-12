package main

import "fmt"

func main() {
	v := []int{1, 2, 6, 3, 9, 4, 7, 5}
	fmt.Println(buscaSoma(v, 9))
	fmt.Println(buscaSoma(v, 4))
	fmt.Println(buscaSoma(v, 19))
}

func buscaSoma(v []int, alvo int) (int, int) {
	n := len(v)

	for i := 0; i <= n - 2; i++ {
		for j := i + 1; j <= n - 1; j++ {
			if v[i] + v[j] == alvo {
				return i, j
			}
		}
	}

	return -1, -1
}