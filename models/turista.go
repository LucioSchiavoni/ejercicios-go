package models

type Turista struct {
	Usuario
}

func (t *Turista) Activo()       { t.activo = true }
func (t *Turista) Bloqueado()    { t.bloqueado = false }
func (t *Turista) Crear() string { return "Manolo" }
func (t *Turista) Rol() string   { return "Turista" }
