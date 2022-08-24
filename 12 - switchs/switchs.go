package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Não definido"
	}
}

func diaDaSemana2(numero int) string {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "Domingo"
	case numero == 2:
		diaSemana = "Segunda"
	case numero == 3:
		diaSemana = "Terça"
	case numero == 4:
		diaSemana = "Quarta"
	case numero == 5:
		diaSemana = "Quinta"
	case numero == 6:
		diaSemana = "Sexta"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Não definido"
	}
	return diaSemana
}

func main() {
	fmt.Println("Switchs")
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(9))

	fmt.Println("Segunda maneira")
	fmt.Println(diaDaSemana2(2))
	fmt.Println(diaDaSemana2(3))
	fmt.Println(diaDaSemana2(4))
	fmt.Println(diaDaSemana2(5))
	fmt.Println(diaDaSemana2(9))
}
