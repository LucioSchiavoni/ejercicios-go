package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func Ejercicio02() string {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero..")
		if input.Scan() {
			numero, err = strconv.Atoi(input.Text())
			if err != nil {
				continue
			} else {
				break
			}

		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}
	return texto
}
