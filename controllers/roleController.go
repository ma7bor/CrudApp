package controllers

import (
	"crud_app/models"
	"crud_app/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct {
	RoleRepository *repositories.RoleRepository
}

func NewRoleController(roleRepo *repositories.RoleRepository) *RoleController {
	return &RoleController{RoleRepository: roleRepo}
}

func (uc *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.RoleRepository.Create(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}
