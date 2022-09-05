package main

import (
	"ex1-aula/exe1"
	"ex1-aula/exe2"
	"ex1-aula/exe3"
	"ex1-aula/exe4"
	"ex1-aula/exe5"
	"ex1-aula/exe6"
	"ex1-aula/exe7"
	"fmt"
	"os"
)

func main() {
	var (
		nome  string
		opcao int
	)
	fmt.Println("Informe o seu nome")
	fmt.Scan(&nome)
	fmt.Println("Olá sr. ", nome)
	for {
		fmt.Println("1 - Encontrar o maior número")
		fmt.Println("2 - Encontrar o maior número")
		fmt.Println("3 - Fatorial")
		fmt.Println("4 - Calculadora")
		fmt.Println("5 - Escolha da profissão")
		fmt.Println("6 - Calcular desconto")
		fmt.Println("7 - Monitoramento")
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
			exe4.Calculadora()
		case 5:
			fmt.Println("Opção 5 escolhida\n")
			exe5.Profissao()
		case 6:
			fmt.Println("Opção 6 escolhida\n")
			exe6.Desconto()
		case 7:
			fmt.Println("Opção 7 escolhida\n")
			exe7.MonitoramentoSites()
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)

		}
	}

}
