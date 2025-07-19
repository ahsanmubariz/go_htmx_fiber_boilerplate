package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `form:"name" validate:"required,min=2"`
	Email string `form:"email" gorm:"unique" validate:"required,email"`
}
