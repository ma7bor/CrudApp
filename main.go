package main

import (
	"crud_app/config"
	"crud_app/initDb"
	"crud_app/repositories"
	"crud_app/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Path to your SQL script file
	sqlScriptPath := "./initDb/initData.sql"

	// Initialize the database
	db, err := initDb.InitializeDatabase(sqlScriptPath)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)
	roleRepository := repositories.NewRoleRepository(db)
	userRolesRepository := repositories.NewUserRolesRepository(db)

	// Initialize Gin router and routes
	r := gin.Default()
	routes.InitializeUserController(r, userRepository, roleRepository, userRolesRepository)

	// Start the server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

//db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/go_test_db?charset=utf8")
//if err != nil {
//	panic(err)
//}
//defer db.Close()
//
//r := routes.SetupRouter(db)
//r.Run(":8880")

//router := gin.Default()
//router.GET("/api/home", handleHome)
//err := router.Run("localhost:9091")
//if err != nil {
//	panic(err)
//}}
