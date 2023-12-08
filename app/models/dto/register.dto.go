package dto

type RegisterDTO struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}
