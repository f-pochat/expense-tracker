package dto

type LoginDTO struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
