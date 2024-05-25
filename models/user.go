package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int    `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstName"` // Use proper casing for JSON fields
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"` // Add validation (email format)
	Tel       string `json:"tel"`
	RoleID    uint   `json:"roleId" gorm:"foreignKey:RoleID;references:ID"` // Foreign key for Role
	Role      Role   `gorm:"foreignKey:RoleID;references:ID"`               // Belongs To relationship with Role
}
