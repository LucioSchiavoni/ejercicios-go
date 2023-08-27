package models

type Usuario struct {
	Id        int
	name      string
	password  string
	email     string
	activo    bool
	bloqueado bool
	rol       string
}

func (u *Usuario) Activo()       { u.activo = true }
func (u *Usuario) Bloqueado()    { u.bloqueado = false }
func (u *Usuario) Crear() string { return "Jose" }
func (u *Usuario) Rol() string   { return "Admin" }
