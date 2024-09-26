package models

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Dsc       string `json:"dsc"`
	CreatedBy int    `json:"created_by" db:"created_by"`
	UpdatedBy int    `json:"updated_by" db:"updated_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
