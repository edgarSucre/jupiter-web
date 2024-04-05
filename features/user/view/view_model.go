package view

type User struct {
	ID      string
	Admin   bool
	Email   string
	Name    string
	TemplID string
}

type UserForm struct {
	Admin          string
	Email          string
	Name           string
	Password       string
	RepeatPassword string
}
