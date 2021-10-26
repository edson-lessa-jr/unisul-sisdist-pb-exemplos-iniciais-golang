package main

import (
	"ex1-aula/exe1"
	"ex1-aula/exe2"
	"ex1-aula/exe3"
	"ex1-aula/exe4"
	"fmt"
	"os"
)

func main() {
	var (
		nome string
		opcao int
	)
	fmt.Println("Informe o seu nome")
	fmt.Scan(&nome)
	fmt.Println("Olá sr. ", nome)
	for {
		fmt.Println("1 - Encontrar o maior número")
		fmt.Println("2 - Encontrar o maior número")
		fmt.Println("3 - Fatorial")
		fmt.Println("4 - Monitoramento")
		fmt.Println("0 - Sair")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Opção 1 escolhida\n")
			exe1.EncontrarMaiorNumero()
		case 2:
			fmt.Println("Opção 2 escolhida\n")
			exe2.CalcularEquacao()
		case 3:
			fmt.Println("Opção 3 escolhida\n")
			exe3.CalcularFatorial()
		case 4:
			fmt.Println("Opção 4 escolhida\n")
			exe4.MonitoramentoSites()
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)

		}
	}

}
