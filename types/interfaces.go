package types

// Solo pasar funciones
type Acciones interface {
	Crear() string
	Activo()
	Bloqueado()
	Rol() string
}
