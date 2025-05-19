package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Wesley",
		"sobrenome": "de Andrade Santos",
		"idade":     "31",
		"peso":      "92",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"nome":      "Lucas",
			"sobrenome": "Silva",
			"idade":     "25",
			"peso":      "80",
		},
		"curso": {
			"nome":      "Tecnologia de Segurança da Informação",
			"faculdade": "Fatec Araraquara",
			"ano":       "2025",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "curso") //deleta o map
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"descricao":  "cancer",
		"nascimento": "05/07/1993",
		"idade":      "31",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["signo"]["nascimento"]) //imprime a data de nascimento
	fmt.Println(usuario2["signo"]["descricao"])  //imprime o signo
}
