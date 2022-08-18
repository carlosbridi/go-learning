package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	endereco string
	numero   uint8
}

func main() {

	fmt.Println("Arquivo structs")

	var usu usuario
	usu.nome = "Bridi"
	usu.idade = 29
	fmt.Println(usu)

	usuario2 := usuario{"Bridi", 29, endereco{"Av.", 9}}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 29}
	fmt.Println(usuario3)
}
