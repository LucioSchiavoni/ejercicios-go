package files

import (
	"bufio"
	"cursoGo/ejercicios"
	"fmt"

	"os"
)

var fileName string = "./files/txt/tabla.txt"

func GuardarTabla() {
	var texto string = ejercicios.Ejercicio02()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close() //Todo archivo que se crea para guardar tiene que cerrarse
}

func SumaTabla() {
	var texto string = ejercicios.Ejercicio02()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido :(")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644) //damos los permisos al usuario con un 0 a la izquierda siempre
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeerArchivo() { //Va a leer el archivo que se creo
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()

}
