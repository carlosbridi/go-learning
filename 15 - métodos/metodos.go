package main

import "fmt"

type usuario struct {
	idade uint8
	name  string
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário: %s no banco de dados\n", u.name)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario := usuario{29, "Bridi"}
	usuario.salvar()
	fmt.Println(usuario.maiorIdade())
	usuario.fazerAniversario()
	fmt.Println(usuario.idade)
}
