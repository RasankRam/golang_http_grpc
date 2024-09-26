package models

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
