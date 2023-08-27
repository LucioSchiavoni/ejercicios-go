package main

import (
	e "cursoGo/ejercicios/interfaces"
	"cursoGo/models"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// estado, respuesta := variables.ConviertoString(5001)
	// fmt.Println(estado)
	// fmt.Println(respuesta)
	// variables.ValidarSo()
	// variables.ValidarSoCase()
	// numero, texto := ejercicios.Ejercicio1("20")
	// fmt.Println(numero)
	// fmt.Println(texto)
	// inputs.IngresoNumeros()
	// iteraciones.Iterar()
	// fmt.Println(ejercicios.Ejercicio02())
	// files.LeerArchivo()
	// funciones.Calculos()
	// funciones.CallClosure()
	// funciones.Exponencia(2)
	// // slices.MuestroArray()
	// variables.ValidarSoCase()
	// slices.MuestroSlice()
	// slices.Capacidad()
	// users.AltaUsuario()
	Manolito := new(models.Usuario)
	e.ActivandoUsuarios(Manolito)
}
