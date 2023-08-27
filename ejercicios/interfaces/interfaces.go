package interfaces

import (
	"cursoGo/types"
	"fmt"
)

// pasa el interfaces de funciones para usar el Rol()
func ActivandoUsuarios(usr types.Acciones) {
	usr.Rol()
	fmt.Printf("Soy rol %s, y soy genial\n", usr.Rol())
}
