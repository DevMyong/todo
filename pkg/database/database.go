package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Database is interface to connect db
type Database interface {
	Connect() error
	Disconnect() error
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=todo_app port=5432 sslmode=disable TimeZone=Asia/Seoul"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 마이그레이션
	DB.AutoMigrate(&Todo{})
}

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
