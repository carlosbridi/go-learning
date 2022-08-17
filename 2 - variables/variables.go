package main

import "fmt"

func main() {
	var variavel1 string = "Olá mundo!"
	fmt.Println(variavel1)

	variavel2 := "Olá mundo 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Carlos"
		variavel4 string = "Bridi"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "abc", "def"
	fmt.Println(variavel5, variavel6)

	const constante1 float64 = 3.141516
	fmt.Println(constante1)

	variavel3, variavel4 = variavel4, variavel3
	fmt.Println(variavel3, variavel4)
}
