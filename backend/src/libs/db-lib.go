package lib

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 生のSQLバージョン
func DBConnectSetUp() (*sql.DB, error) {
	fmt.Println("Connecting to database...")

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// 接続文字列を修正
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// デバッグ用に接続情報を出力
	fmt.Printf("Debug - Connection params:\nHost: %s\nPort: %s\nUser: %s\nDB: %s\n",
		host, port, user, dbname)

	// データベース接続
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// データベースに接続
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping database:", err)
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Connected to database")
	return db, nil
}

// GORMバージョン
func DBConnectSetUpGORM() (*gorm.DB, error) {
	fmt.Println("GORM - Connecting to database...")

	// 接続文字列を修正
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	// デバッグ用に接続情報を出力
	fmt.Printf("Debug - Connection params:\nHost: %s\nPort: %s\nUser: %s\nDB: %s\n",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"))

	// GORMの接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("GORM - Failed to connect to database:", err)
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	fmt.Println("GORM - Connected to database")
	return db, nil
}
