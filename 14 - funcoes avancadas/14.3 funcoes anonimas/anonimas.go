package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		fmt.Println(texto)
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Ola mundo!")

	fmt.Println(retorno)
}
