package variables

import (
	"fmt"
	"runtime"
)

func ValidarSo() {

	if so := runtime.GOOS; so == "Windows." {
		fmt.Println("Esto es windows")
	} else if so == "linux." {
		fmt.Println("Sistema Linux")
	} else if so == "darwin" {
		fmt.Println("Esto es mac!!! wow")
	} else {
		fmt.Println("Sistema operativo desconocido")
	}

}

func ValidarSoCase() {

	switch so := runtime.GOOS; so {
	case "windows":
		fmt.Println("Estas en windows")

	case "darwin":
		fmt.Println("Estas en Mac")
	case "linux":
		fmt.Println("Estas en Linux")
	default:
		fmt.Printf("Estas en %s \n", so) //devuelve el sistema que esta usando
	}
}
