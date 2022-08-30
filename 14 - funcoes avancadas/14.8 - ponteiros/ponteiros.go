package main

import "fmt"

func inverterSinal(n1 int) int {
	return n1 * -1
}

//Alterando valor do endereço de memória
func inverterSinalComPonteiro(n1 *int) {
	*n1 = *n1 * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)

	//Manipulação usando valor do ponteiro
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
