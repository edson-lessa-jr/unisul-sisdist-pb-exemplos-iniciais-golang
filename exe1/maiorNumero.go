package exe1

import "fmt"

func EncontrarMaiorNumero() {

	var (
		notas [5]float64
		maior float64
	)

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Infome a %d nota\n",i)
		var nota float64
		fmt.Scan(&nota)
		notas[i]=nota
		if (i==0) || (maior<nota){
			maior=nota
		}

	}
	fmt.Println("Dado as notas: ", notas)
	fmt.Println("A maior nota é: ", maior)

}