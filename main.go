package main

import (
	"cursoGo/middleware"
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
	// Manolito := new(models.Usuario)
	// e.ActivandoUsuarios(Manolito)
	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()
	// canal1 := make(chan bool)
	// // go para decir que es async
	// go goroutines.MiNombreLento("Lucio Schiavoni", canal1)
	// //es como un await en node y defer para dejarlo a lo ultimo
	// defer func() {
	// 	<-canal1
	// }()
	// fmt.Println("Bienvenido")
	// webserver.ServidorWeb()
	middleware.MiMiddleware()
}
