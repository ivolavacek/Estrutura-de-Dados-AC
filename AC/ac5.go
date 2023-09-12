package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 7, 9, 11, 14}
	fmt.Println(buscaSoma(v, 9))
	fmt.Println(buscaSoma(v, 4))
	fmt.Println(buscaSoma(v, 19))
	fmt.Println(buscaSoma(v, 20))
}

func buscaSoma(v []int, alvo int) (int, int) {
	var inf = 0
	var sup = len(v) - 1

	for inf < sup {
		var soma = v[inf] + v[sup]

		if soma == alvo {
			return v[inf], v[sup]
		} else if v[sup] >= alvo {
			sup --
			fmt.Println("sup", v[sup])
		} else {
			inf ++
			fmt.Println("inf", v[inf])
		}
	}

	return -1, -1
}