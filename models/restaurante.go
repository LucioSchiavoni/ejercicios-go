package models

type Retaurante struct {
	Usuario
}

func (r *Retaurante) Activo()       { r.activo = true }
func (r *Retaurante) Bloqueado()    { r.bloqueado = false }
func (r *Retaurante) Crear() string { return "JuanJose" }
func (r *Retaurante) Rol() string   { return "Restaurante" }
