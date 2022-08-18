package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var numero1 int8 = 10
	var numero2 int8 = numero1

	fmt.Println(numero1, numero2)

	var numero3 int8 = 10
	var numero4 *int8

	fmt.Println(numero3, numero4)

	var numero5 int8 = 12
	var ponteiro *int8

	ponteiro = &numero5
	fmt.Println(numero5, ponteiro)

	//Desreferenciação
	fmt.Println(numero5, *ponteiro)
	numero5++
	fmt.Println("Pós desreferenciação")
	fmt.Println(numero5, *ponteiro)

}
