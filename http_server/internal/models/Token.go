package models

type Token struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	Token     string `json:"token"`
	UserAgent string `json:"user_agent" db:"user_agent"`
	Ip        string `json:"ip"`
}
