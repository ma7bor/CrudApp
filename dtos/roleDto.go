package dtos

import "crud_app/models"

type RoleDto struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Users       []models.User `json:"users"`
}
