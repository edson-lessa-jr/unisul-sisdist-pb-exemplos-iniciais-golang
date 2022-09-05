package exe5

import (
	"fmt"
	"strings"
)

func Profissao() {
	var (
		opcao string
	)
menu:
	for {
		fmt.Println("Escolha a profissão:\n" +
			"e - engenheiro\n" +
			"m - médico\n" +
			"d - designer\n" +
			"p - programador\n" +
			"a - advogado\n" +
			"v - voltar")
		fmt.Scan(&opcao)

		opcao = strings.ToLower(opcao)
		switch opcao {
		case "e":
			fmt.Println("Tibúrcio é engenheiro")
		case "m":
			fmt.Println("Tibúrcio é médico")
		case "d":
			fmt.Println("Tibúrcio é designer")
		case "p":
			fmt.Println("Tibúrcio é programador")
		case "a":
			fmt.Println("Tibúrcio é advogado")
		default:
			fmt.Println("Tibúrcio está desempregado")
		}
		fmt.Println("\n\nDeseja voltar? S/N\n\n")
		fmt.Scan(&opcao)
		opcao = strings.ToLower(opcao)
		if opcao == "S" {
			break menu
		}
	}
}
