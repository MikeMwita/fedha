package postgres

import (
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// FOR GORM
var db *gorm.DB
var err error

//type User struct {
//	gorm.Model
//	UserId      string
//	UserName    string
//	Email       string
//	PhoneNumber string
//	Hash        string
//}

func NewPostgresDB() (*gorm.DB, error) {
	dbHost := os.Getenv("POSTGRES_DB_HOST")
	dbPort := os.Getenv("POSTGRES_DB_PORT")
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PSWD")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	fmt.Println("Connection String:", connStr)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//migrate the schema
	db.AutoMigrate(models.User{}, models.Expense{}, models.ExpenseType{})
	if err != nil {
		log.Fatalf("Error applying migrations: %s", err)
	}

	return nil, err
}
