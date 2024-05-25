package repositories

import (
	"crud_app/models"
	"errors"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (ur *RoleRepository) Create(role *models.Role) error {
	return ur.DB.Create(role).Error
}

func (r *RoleRepository) Exists(roleID uint) (bool, error) {
	var role models.Role
	err := r.DB.First(&role, roleID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // Role does not exist
		}
		return false, err // Error occurred
	}
	return true, nil // Role exists
}
