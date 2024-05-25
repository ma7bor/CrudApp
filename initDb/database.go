package initDb

import (
	"crud_app/config"
	"crud_app/models"
	"fmt"
	"github.com/goccy/go-json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// ConnectToMySQL connects to the MySQL database using the given DSN
func ConnectToMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db, err
}

// CreateDatabaseIfNotExists creates the database if it does not already exist
func CreateDatabaseIfNotExists(db *gorm.DB, dbName string) error {
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	return db.Exec(createDBSQL).Error
}

// MigrateTables runs the AutoMigrate function to create tables based on the models
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{})
}

// InitializeDatabase initializes the database, ensuring it exists and is migrated
func InitializeDatabase(scriptPath string) (*gorm.DB, error) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Connect to MySQL server without specifying the database
	rootDSN := "root:root@tcp(localhost:3306)/?parseTime=true"
	rootDB, err := ConnectToMySQL(rootDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL server: %w", err)
	}

	// Ensure the database exists
	dbName := cfg.DBName
	if err := CreateDatabaseIfNotExists(rootDB, dbName); err != nil {
		return nil, fmt.Errorf("failed to create database %s: %w", dbName, err)
	}

	// Connect to the specific database
	dsn := fmt.Sprintf("root:root@tcp(localhost:3306)/%s?parseTime=true", dbName)
	db, err := ConnectToMySQL(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Run migrations
	if err := MigrateTables(db); err != nil {
		return nil, fmt.Errorf("failed to migrate tables: %w", err)
	}

	// Execute SQL script to initialize the database with data
	if err := ExecuteSQLScript(db, scriptPath); err != nil {
		return nil, fmt.Errorf("failed to execute SQL script: %w", err)
	}

	var countUsers int64
	err = db.Table("users").Count(&countUsers).Error
	// Insert users
	if countUsers == 0 {
		if err := InsertUsers(db); err != nil {
			return nil, fmt.Errorf("failed to insert users: %w", err)
		}
	}

	var countUserRoles int64
	err = db.Table("users").Count(&countUserRoles).Error
	if countUserRoles == 0 {
		// Read user roles from JSON file
		userRoles, err := ReadUserRolesFromFile("./initDb/userRolesData.json")
		if err != nil {
			log.Fatalf("failed to read user roles from file: %v", err)
		}
		// Insert user roles into the database
		err = InsertUserRoles(db, userRoles)
		if err != nil {
			log.Fatalf("failed to insert user roles into the database: %v", err)
		}
	}

	return db, nil
}

// ExecuteSQLScript executes a SQL script to initialize the database with data
func ExecuteSQLScript(db *gorm.DB, scriptPath string) error {
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to read SQL script file: %w", err)
	}
	statements := string(script)
	return db.Exec(statements).Error
}

func InsertUsers(db *gorm.DB) error {
	users := []models.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Tel: "1234567890", RoleID: 1},
		{ID: 2, FirstName: "Alice", LastName: "Smith", Email: "alice@example.com", Tel: "9876543210", RoleID: 2},
		{ID: 3, FirstName: "Bob", LastName: "Johnson", Email: "bob@example.com", Tel: "5556667777", RoleID: 2},
		{ID: 4, FirstName: "Eva", LastName: "Brown", Email: "eva@example.com", Tel: "9998887777", RoleID: 3},
		{ID: 5, FirstName: "Michael", LastName: "Clark", Email: "michael@example.com", Tel: "1112223333", RoleID: 3},
		{ID: 6, FirstName: "Sarah", LastName: "Wilson", Email: "sarah@example.com", Tel: "4445556666", RoleID: 3},
		{ID: 7, FirstName: "David", LastName: "Martinez", Email: "david@example.com", Tel: "7778889999", RoleID: 3},
		{ID: 8, FirstName: "Jennifer", LastName: "Garcia", Email: "jennifer@example.com", Tel: "2223334444", RoleID: 3},
	}

	return db.Create(&users).Error
}

func ReadUserRolesFromFile(filePath string) ([]models.UserRole, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("failed to close json access file: %w", err)
		}
	}(file)

	// Decode the JSON data into a slice of user roles
	var userRoles []models.UserRole
	err = json.NewDecoder(file).Decode(&userRoles)
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

func InsertUserRoles(db *gorm.DB, userRoles []models.UserRole) error {
	for _, userRole := range userRoles {
		// Insert each user role into the user_roles table
		if err := db.Create(&userRole).Error; err != nil {
			return err
		}
	}
	return nil
}
