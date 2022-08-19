package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int8
}

func main() {
	fmt.Println("Bridi")

	var array1 [5]int
	fmt.Println(array1)

	var personArray [5]person
	personArray[0].name = "Bridi"
	personArray[0].age = 29
	fmt.Println(personArray)

	array2 := [5]string{"Bridi", "Carlos", "Timbó", "SC", "Brazil"}
	fmt.Println(array2)

	//não é muito utlizado e é sempre necessário passar a quantidade de tamanhos nos arrays
	//os 3 pontos são inferência do tamanho
	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	//Slice
	//Não é um array, mas encorpora um array
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	//adição de dado
	slice = append(slice, 9)
	fmt.Println(slice)

}
