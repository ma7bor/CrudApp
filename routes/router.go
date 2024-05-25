package routes

import (
	"crud_app/controllers"
	"crud_app/repositories"
	"github.com/gin-gonic/gin"
)

var baseUrl string = "/api/users"
var apiUrl string = "/api/users/:id"

func InitializeUserController(router *gin.Engine, userRepository *repositories.UserRepository,
	roleRepository *repositories.RoleRepository,
	userRolesRepository *repositories.UserRolesRepository) {
	userController := controllers.NewUserController(userRepository, roleRepository, userRolesRepository)
	roleController := controllers.NewRoleController(roleRepository)

	// Define routes for user-related operations
	router.POST(baseUrl, userController.CreateUser)
	router.GET(apiUrl, userController.GetUserByID)
	router.GET(baseUrl, userController.GetUsers)
	router.PUT(apiUrl, userController.UpdateUser)
	router.DELETE(apiUrl, userController.DeleteUser)

	//routes for role operations
	router.POST("api/roles", roleController.CreateRole)

}
