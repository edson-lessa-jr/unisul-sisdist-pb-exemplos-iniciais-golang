package exe2

import (
	"fmt"
	"math"
)

func CalcularEquacao()  {

	fmt.Println("Dado uma equação do segundo grau em ax²+bx+c informe as variáveis a, b, c")
	a, b, c := coletarVariavies()
	delta := math.Pow(float64(b), 2) - float64(4*a*c)
	if delta<0{
		fmt.Println("Não é possível calcular a raíz")
		return
	}

	x1, x2 := calcularRaizes(a,b, delta)
	fmt.Printf("As Raízes são x'=%f, x''=%f\n", x1, x2)

}

func calcularRaizes(a int, b int, delta float64) (float64, float64) {
	x1 := (float64(-b) + math.Sqrt(delta))/float64(2.0*a)
	x2 := (float64(-b) - math.Sqrt(delta))/float64(2.0*a)
	return x1, x2

}
func coletarVariavies() (a, b, c int) {
	fmt.Println("Informe a variável a")
	fmt.Scan(&a)
	fmt.Println("Informe a variável b")
	fmt.Scan(&b)
	fmt.Println("Informe a variável c")
	fmt.Scan(&c)
	return

}
