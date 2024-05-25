package repositories

import (
	"crud_app/models"
	"gorm.io/gorm"
)

type UserRolesRepository struct {
	DB *gorm.DB
}

func NewUserRolesRepository(db *gorm.DB) *UserRolesRepository {
	return &UserRolesRepository{DB: db}
}

func (r *UserRolesRepository) Exists(userID, roleID uint) (bool, error) {
	var count int64
	if err := r.DB.Model(&models.UserRole{}).Where("user_id = ? AND role_id = ?", userID, roleID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *UserRolesRepository) Create(userRole *models.UserRole) error {
	return ur.DB.Create(userRole).Error
}

func (r *UserRolesRepository) DeleteByUserID(userID string) error {
	// Delete user roles associated with the user ID
	if err := r.DB.Where("user_id = ?", userID).Delete(&models.UserRole{}).Error; err != nil {
		return err
	}
	return nil
}
