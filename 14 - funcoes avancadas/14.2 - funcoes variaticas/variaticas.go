package main

import "fmt"

func soma(numeros ...int) {
	fmt.Println(numeros)
	var total int = 0

	for i := 0; i < len(numeros); i++ {
		total += numeros[i]
	}

	var total2 int = 0
	for _, numero := range numeros {
		total2 += numero
	}

	fmt.Println(total)
	fmt.Println(total2)
}

func main() {
	fmt.Println("Variaticas")
	soma(1, 2, 3, 4, 5)
}
