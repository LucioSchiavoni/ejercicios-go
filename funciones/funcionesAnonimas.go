package funciones

import "fmt"

// se puede pasar otra func al lado de func Calculos() con dos parametros por ej dos int
// y devolver un numero como resultado para interactuar con el usuario
func Calculos() {

	var numero3 int = 57
	var numero4 int = 64

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(20, 8))

	calculo = func(num1 int, num2 int) int {
		return num1 * num2 * numero3 * numero4
	}

	fmt.Println(calculo(96, 36))
}
