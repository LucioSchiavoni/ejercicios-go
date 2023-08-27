package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Jose"
	Estado = true
	Sueldo = 23.47
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Print(Fecha)

}

// la req es envio que trae un int y se pasa a string
func ConviertoString(envio int) (bool, string) {
	respuesta := strconv.Itoa(envio)
	return true, respuesta
}
