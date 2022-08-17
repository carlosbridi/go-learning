package main

import (
	"errors"
	"fmt"
)

func main() {

	//4 tipos de numeros inteiros no go
	//int8, int16, int32, int64
	//int - usará a arquitetura do PC como base, 64 bits ou 32 bits
	var numero int8 = 100
	fmt.Println(numero)

	//unsigned data type, only > 0
	var numero2 uint8 = 0 //-1;
	fmt.Println(numero2)

	//alias
	//INT32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numberoReal1 float32 = 1.2
	fmt.Println(numberoReal1)

	var numberoReal2 float64 = 3.4
	fmt.Println(numberoReal2)

	//Strings
	var texto1 string = "TextoBridi"
	fmt.Println(texto1)

	charSample := 'A'
	fmt.Println(charSample)
	//End-Strings

	var booleanType bool //default false
	fmt.Println(booleanType)

	var erro error //default nil
	fmt.Println(erro)

	var erro2 error = errors.New("Houve um erro")
	fmt.Println(erro2)
}
