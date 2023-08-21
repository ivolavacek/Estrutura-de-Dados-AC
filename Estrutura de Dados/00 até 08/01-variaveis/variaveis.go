package main

import "fmt"

var w float64

const Pi = 3.14

func main() {
	// declaração explícita
	var x, y int
	var texto string

	x, y = 4, 5
	texto = "olá, mundo"

	fmt.Println(x, y, texto)

	// declaração implícita
	var z = 4.54

	fmt.Println(z)

	// declaração curta
	msg := "olá"

	fmt.Println(msg)

	var num complex128

	num = 48 + 7.58i

	fmt.Println(num)
}

/*
inteiros

int8     -128 a 127
int16    -32... a 32...
int32
int64
uint8     0 a 255
uint16    0 a 65.535
uint32
uint64

int       32bits ou 64bits, depende da arquitetura do pc
byte      uint8
rune      int32  -> caracteres unicode

-----------------
flututantes

float32
float64 -> declaração implícita ou curta

-----------------
complexos

complex64
complex128

-----------------
demais

string
nil
*/
