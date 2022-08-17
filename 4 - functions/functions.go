package main

import "fmt"

func somarValores(valor1 int8, valor2 int8) int8 {
	return valor1 + valor2
}

func resultadoDoisValores(n1, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2
}

func main() {
	soma := somarValores(1, 2)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	var funcao = f("Texto aqui")
	fmt.Println(funcao)

	soma, subtracao := resultadoDoisValores(10, 9)
	fmt.Println(soma, subtracao)

	//sinal de underline ignora segundo parametro
	soma2, _ := resultadoDoisValores(10, 9)
	fmt.Println(soma2)
}
