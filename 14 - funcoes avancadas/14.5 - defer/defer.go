package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoAprovado(nota int, nota2 int) bool {
	defer fmt.Println("Médica calculada, resultado sera calculado")

	fmt.Println("Entrando na funcao de aluno aprovado")
	media := (nota + nota2) / 2

	return media >= 6
}

func main() {
	//Defer adia a execução por último, imediatamente antes do return
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoAprovado(4, 5))
}
