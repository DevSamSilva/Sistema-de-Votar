package funcoes

import (
	"fmt"

	"github.com/DevSamSilva/Sistema-de-Votar/estrutura"
)

var Candidatos []estrutura.Candidato

func AdicionarCandidato(nome string, numero int64) {

	var qtd int
	fmt.Println("Quantos candidatos?")
	fmt.Scan(&qtd)

	for i := 1; i <= qtd; i++ {
		fmt.Println("Digite o nome do candidato(a)")
		fmt.Scan(&nome)
		fmt.Println("Digite o numero do canditato")
		fmt.Scan(&numero)

		CandidatoNovo := estrutura.Candidato{
			Nome:   nome,
			Numero: numero,
		}

		Candidatos = append(Candidatos, CandidatoNovo)
	}

}

func MostrarCandidatos() {
	for _, c := range Candidatos {
		fmt.Printf("CANDIDATO(A): %s\nNÚMERO: %d\nVOTOS: %d\n****************\n", c.Nome, c.Numero, c.VotosTotal)
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
