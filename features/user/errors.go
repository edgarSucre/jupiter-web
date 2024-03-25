package user

import "errors"

var ErrNoName = errors.New("nombre es requerido")
var ErrNoUserName = errors.New("nombre de usuario es requerido")
var ErrNoPassword = errors.New("password es requerido")
var ErrPasswordNoMatch = errors.New("passwords no son iguales")
