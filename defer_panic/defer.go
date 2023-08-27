package defer_panic

//Cuando termine la funcion
//se ejecuta lo que se puso en defer

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Aqui termina todo")

	fmt.Println("Segundo mensaje")
}

func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Se encontro un error ,%v \n", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Error de panico")
	}
}
