package domain

import "errors"

type MyErr struct {
	msg string
}

func (me MyErr) Error() string {
	return me.msg
}

func NewErr(msg string) MyErr {
	return MyErr{
		msg: msg,
	}
}

var (
	ErrNotFound = NewErr("no se encotro entidad en la bd")
)

func ViewErr(err error, defaultMsg string) map[string]string {
	if errors.As(err, new(MyErr)) {
		return map[string]string{"title": err.Error()}
	}

	return map[string]string{"title": defaultMsg}
}
