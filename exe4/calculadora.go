package exe4

import (
	"fmt"
	"strings"
)

func Calculadora() {
	var (
		opcao            string
		numero1, numero2 float64
	)
menu:
	for {

		fmt.Println("Informe o primeiro número")
		fmt.Scan(&numero1)
		fmt.Println("Informe o segundo número")
		fmt.Scan(&numero2)
		fmt.Println("+ - Adição")
		fmt.Println("- - Subtração")
		fmt.Println("* - Multiplicação")
		fmt.Println("/ - Divisão")
		fmt.Scan(&opcao)
		opcao = strings.ToLower(opcao)

		switch opcao {
		case "+":
			fmt.Printf("Resultado da adição é %f", numero1+numero2)
		case "-":
			fmt.Printf("Resultado da subtração é %f", numero1-numero2)
		case "*":
			fmt.Printf("Resultado da multiplicação é %f", numero1*numero2)
		case "/":
			fmt.Printf("Resultado da divisão é %f", numero1/numero2)
		case "s":
			break menu
		}
		fmt.Println("\n\nDeseja voltar? S/N\n\n")
		fmt.Scan(&opcao)
		opcao = strings.ToLower(opcao)
		if opcao == "S" {
			break menu
		}
	}
}
