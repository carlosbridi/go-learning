package main

import "fmt"

func recuperarExec() {
	fmt.Println("Recuperando")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println("Panic")
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execucação")
}
