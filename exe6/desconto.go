package exe6

import (
	"fmt"
	"strings"
)

func Desconto() {
	var (
		nome, opcao       string
		preco, precoFinal float64
	)
menu:
	for {
		fmt.Println("Informe o nome do cliente")
		fmt.Scan(&nome)
		fmt.Println("Informe o preço")
		fmt.Scan(&preco)
		fmt.Println("Escolha a categoria:\n" +
			"A - 10%\n" +
			"B - 15%\n" +
			"C - 20%\n" +
			"D - 25%\n" +
			"E - 50%\n")
		fmt.Scan(&opcao)
		opcao = strings.ToUpper(opcao)
		switch opcao {
		case "A":
			precoFinal = preco * (1 - 0.10)
		case "B":
			precoFinal = preco * 0.85
		case "C":
			precoFinal = preco * 0.80
		case "D":
			precoFinal = preco * 0.75
		case "E":
			precoFinal = preco * 0.50
		default:
			fmt.Println("Sem desconto")
		}
		if precoFinal > 0 {
			fmt.Printf("Cliente: %s\nPreço: %f\nPreço final: %f\n", nome, preco, precoFinal)
		}
		fmt.Println("\n\nDeseja voltar? S/N\n\n")
		fmt.Scan(&opcao)
		opcao = strings.ToLower(opcao)
		if opcao == "S" {
			break menu
		}
	}
}
