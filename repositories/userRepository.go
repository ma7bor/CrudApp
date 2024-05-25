package repositories

import (
	"crud_app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := ur.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.DB.Delete(user).Error
}
