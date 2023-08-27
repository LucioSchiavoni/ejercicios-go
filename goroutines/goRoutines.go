package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") //split separa los caracteres junto con un separador
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //duracion
		fmt.Println(letra)
	}
	//forma de asignar un chan <-
	canal1 <- true
}
