package main

import (
	"fmt"

	"github.com/DevSamSilva/Sistema-de-Votar/funcoes"
)

func main() {

	var nome string
	var numero int64

	for {
		var opcoes int
		fmt.Println("Selecione: (1)Adicionar Candidato | (2)Mostrar Candidatos | (3)Votar | (4)Resultado")
		fmt.Scan(&opcoes)
		if opcoes <= 0 || opcoes > 4 {
			fmt.Println("Digitos invalidos")
			continue
		}

		switch opcoes {
		case 1:
			funcoes.AdicionarCandidato(nome, numero)
		case 2:
			funcoes.MostrarCandidatos()
		case 3:
			funcoes.Votar(numero)
		case 4:
			funcoes.Resultado()
			return

		}

	}

}
