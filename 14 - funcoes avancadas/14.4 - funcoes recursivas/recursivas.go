package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	fmt.Println("Função recursiva")

	posicao := uint(12)

	for i := 1; i <= int(posicao); i++ {
		fmt.Println(fibonacci(uint(i)))
	}

}
