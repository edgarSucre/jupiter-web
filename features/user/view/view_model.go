package view

type User struct {
	Admin   bool
	TemplID string
	Name    string
}

type UserForm struct {
	Name           string
	UserName       string
	Password       string
	RepeatPassword string
}
