package users

import (
	"cursoGo/models"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(models.User) //nueva instancia de usuario
	u.AddUser(2, "Manolo", time.Now(), true)

	//se agrega usuario ingresando datos como en una request
	fmt.Print(u)
}
