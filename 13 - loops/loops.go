package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// var i int = 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando ", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice)
		fmt.Println(nome)
	}

	fmt.Println("Segundo método de array")
	for _, nome := range nomes {
		fmt.Println(nome)

	}

	fmt.Println(nomes)

	for indice, nome := range "PALAVRA" {
		fmt.Println(indice, string(nome))
	}

	var texto string = "PALAVRA"
	for j := len(texto); j > 0; j-- {
		fmt.Println(j)
	}

}
