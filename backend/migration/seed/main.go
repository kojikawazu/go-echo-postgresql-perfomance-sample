package main

import (
	"fmt"
	"log"
	"os"

	seeds "backend/migration/seed/seeds"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// シードデータを挿入
func main() {
	fmt.Println("Seeding started")

	// .envファイルの読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// データベース接続
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	// GORMの接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// シードデータの挿入
	if err := seeds.SeedAll(db); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	fmt.Println("Seeding completed successfully")
}
