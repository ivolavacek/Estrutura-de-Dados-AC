package main

import "fmt"

/*
t -> T1 -> T2 -> T3 -> T4

PROGRAMA TOMADAS
	Lê T1, T2, T3, T4
	totalTomadas := T1 + T2 + T3 + T4
	Escreve totalTomadas - 3
*/
func main() {
	var t1, t2, t3, t4 int

	fmt.Scanln(&t1, &t2, &t3, &t4)
	totalTomadas := t1 + t2 + t3 + t4

	fmt.Println(totalTomadas - 3)
}