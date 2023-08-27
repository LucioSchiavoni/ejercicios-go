package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := OperacionMiddleware(sumar)(18, 18)
	fmt.Println(result)
	result = OperacionMiddleware(restar)(24, 7)
	fmt.Println(result)
	result = OperacionMiddleware(multiplicar)(7, 3)
	fmt.Println(result)
	result = OperacionMiddleware(dividir)(1020, 5)
	fmt.Println(result)
}

func OperacionMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Print("Inicio de Operacion\n")
		return f(a, b)
	}

}
