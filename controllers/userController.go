package controllers

import (
	"crud_app/models"
	"crud_app/repositories"
	"crud_app/responseDtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

type UserController struct {
	UserRepository      *repositories.UserRepository
	RoleRepository      *repositories.RoleRepository
	UserRolesRepository *repositories.UserRolesRepository
}

func NewUserController(userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository, userRolesRepo *repositories.UserRolesRepository) *UserController {
	return &UserController{
		UserRepository:      userRepo,
		RoleRepository:      roleRepo,
		UserRolesRepository: userRolesRepo,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if all required properties are present
	if err := uc.validateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserRepository.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Use mapstructure to decode the user struct into the custom response struct
	var response responseDtos.UserResponse
	if err := mapstructure.Decode(user, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.UserRepository.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Use mapstructure to decode the user struct into the custom response struct
	var response responseDtos.UserResponse
	if err := mapstructure.Decode(user, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Create a slice to store the transformed user data
	var usersResponse []responseDtos.UserResponse

	// Transform each user entity into the custom response format
	for _, user := range users {
		var userResponse responseDtos.UserResponse
		if err := mapstructure.Decode(user, &userResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usersResponse = append(usersResponse, userResponse)
	}
	c.JSON(http.StatusOK, usersResponse)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Handle the error if the ID parameter is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the role ID sent in the update request
	if user.RoleID != 0 {
		if err := uc.validateRoleID(user.RoleID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Call the method to create user role if not exist
	if err := uc.CreateUserAndRoleIfNotExist(uint(userID), user.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserRepository.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Use mapstructure to decode the user struct into the custom response struct
	var response responseDtos.UserResponse
	if err := mapstructure.Decode(user, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.UserRepository.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := uc.UserRepository.Delete(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete the user roles associated with the user
	if err := uc.UserRolesRepository.DeleteByUserID(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Function to validate if all required properties are present
func (uc *UserController) validateUser(user models.User) error {
	if user.FirstName == "" {
		return errors.New("missing required field: firstName")
	}
	if user.LastName == "" {
		return errors.New("missing required field: lastName")
	}
	if user.Email == "" {
		return errors.New("missing required field: email")
	}
	if user.RoleID == 0 {
		return errors.New("missing required field: roleId")
	}
	return nil
}

func (uc *UserController) validateRoleID(roleID uint) error {
	// Check if the role ID exists in the roles table
	roleExists, err := uc.RoleRepository.Exists(roleID)
	if err != nil {
		return err
	}
	if !roleExists {
		return errors.New("Role does not exist")
	}
	return nil
}

func (uc *UserController) CreateUserAndRoleIfNotExist(userID, roleID uint) error {
	// Check if there's already an entry for the user in the user_roles table
	userRoleExists, err := uc.UserRolesRepository.Exists(userID, roleID)
	if err != nil {
		return err
	}

	// If there's no entry for the user in the user_roles table, add a new entry
	if !userRoleExists {
		userRole := models.UserRole{
			UserID: userID,
			RoleID: roleID,
		}
		if err := uc.UserRolesRepository.Create(&userRole); err != nil {
			return err
		}
	}

	return nil
}
