package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//Dentro do colchete é a chave, fora é o valor
	usuario := map[string]string{
		"nome":      "Carlos",
		"sobrenome": "Bridi",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "João",
			"sobrenome": "Pedro",
		},
		"curso": {
			"nome":   "SI",
			"campus": "Uniasselvi",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}
