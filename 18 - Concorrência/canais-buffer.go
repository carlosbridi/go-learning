package main

import "fmt"

func main() {
	//canal com buffer somente bloqueia quando chegar no limite do canal
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Olá mundo!!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
