package ejercicios

import (
	"strconv"
)

func Ejercicio1(request string) (int, string) {

	respuesta, err := strconv.Atoi(request)

	if err != nil {
		return 0, "Error" + err.Error()
	}
	if respuesta > 100 {
		return respuesta, "Es es mayor a 100"
	} else {
		return respuesta, "Es es menor a 100"
	}

}
