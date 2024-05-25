package dtos

import "crud_app/models"

type UserDto struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Tel       string      `json:"tel"`
	Role      models.Role `gorm:"RoleID"`
}
