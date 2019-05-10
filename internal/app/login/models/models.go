package models

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}
