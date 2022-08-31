package main

import "fmt"

var n int

func init() {
	fmt.Println("Execução da função init")
	n = 10
}

func main() {
	fmt.Println("Função main")
	fmt.Println(n)
}
