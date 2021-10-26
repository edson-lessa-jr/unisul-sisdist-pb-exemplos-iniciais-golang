package exe3

import "fmt"

func CalcularFatorial() {
	fmt.Println("Informe o número para calcular o fatorial")
	var numero int
	fmt.Scan(&numero)
	fmt.Printf("O Fatorial de %d é: %d\n\n", numero, fatorial(numero))

}
func fatorial(numero int) int {

	if numero==1{
		fatorial:= 1
		return fatorial
	} else {
		fatorial:=numero*fatorial(numero-1)
		return fatorial
	}

}