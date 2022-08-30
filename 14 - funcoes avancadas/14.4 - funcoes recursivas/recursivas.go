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

	fmt.Println(fibonacci(10))

	for i := 1; i <= 7; i++ {
		fmt.Println(fibonacci(uint(i)))
	}

}
