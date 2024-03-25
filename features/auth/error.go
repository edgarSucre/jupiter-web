package auth

import (
	"github.com/edgarSucre/jw/domain"
)

var ErrBadCredentials = domain.NewErr("usuario o contrasenña incorrectos")
