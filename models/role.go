package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          int    `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" validate:"required"` // Add validation with a library like "github.com/go-playground/validator"
	Description string `json:"description"`
	Users       []User `gorm:"many2many:user_roles"` // Relationship with User (many-to-many)
}
