package main

import (
	"fmt"
)

type Candidato struct {
	Nome       string
	Numero     int64
	VotosTotal int64
}

var Candidatos []Candidato

func AdicionarCandidato(nome string, numero int64) {

	var qtd int
	fmt.Println("Quantos candidatos?")
	fmt.Scan(&qtd)

	for i := 1; i <= qtd; i++ {
		fmt.Println("Digite o nome do candidato(a)")
		fmt.Scan(&nome)
		fmt.Println("Digite o numero do canditato")
		fmt.Scan(&numero)

		CandidatoNovo := Candidato{
			Nome:   nome,
			Numero: numero,
		}

		Candidatos = append(Candidatos, CandidatoNovo)
	}

}

func MostrarCandidatos() {
	for _, c := range Candidatos {
		fmt.Printf("Candidato: %s\n Numero: %d\n Votos: %d\n", c.Nome, c.Numero, c.VotosTotal)
	}
}

func Votar(numero int64) {

	fmt.Println("******VOTAÇÃO******")
	fmt.Println("Digite o numero do seu candidato: ")
	fmt.Scan(&numero)

	for i := range Candidatos {
		if Candidatos[i].Numero == numero {
			Candidatos[i].VotosTotal++
			fmt.Println("Voto Confirmado!")
			return
		}
	}

	fmt.Println("Candidato não encontrado")
}

func Resultado() {
	fmt.Printf("****RESULTADO FINAL****")
	for _, c := range Candidatos {
		fmt.Printf("Candidato(a): %s\nNumero de Votos: %d\n", c.Nome, c.VotosTotal)
	}
}

func main() {

	var nome string
	var numero int64

	for {
		var opcoes int
		fmt.Println("Selecione: (1)Adicionar Candidato | (2)Mostrar Candidatos | (3)Votar | (4)Resultado")
		fmt.Scan(&opcoes)

		switch opcoes {
		case 1:
			AdicionarCandidato(nome, numero)
		case 2:
			MostrarCandidatos()
		case 3:
			Votar(numero)
		case 4:
			Resultado()
			return

		}

	}

}
