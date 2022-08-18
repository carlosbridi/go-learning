package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //traz os campos da struct para dentro da definição
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	var estudante estudante
	estudante.nome = "Carlos"
	estudante.sobrenome = "Bridi"
	estudante.idade = 29
	estudante.altura = 178
	estudante.curso = "Sistemas"
	estudante.faculdade = "Uniasselvi"
	fmt.Println(estudante)
}
