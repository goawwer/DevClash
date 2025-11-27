package dto

type SignUpForm struct {
	Email    string
	Username string
	Name     string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}
