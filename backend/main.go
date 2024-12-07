package main

import (
	"fmt"
	"os"

	lib "backend/src/libs"
	"backend/src/middleware"
	routes "backend/src/routes"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// PostgreSQL接続設定
	db, err := lib.DBConnectSetUp()
	if err != nil {
		return
	}
	defer db.Close()

	// Gorm接続設定
	dbGorm, err := lib.DBConnectSetUpGORM()
	if err != nil {
		return
	}

	// Echoインスタンス作成
	e, err := lib.EchoSetUp()
	if err != nil {
		return
	}

	// ミドルウェアの設定
	middleware.SetUpMiddleware(e)

	// ルートエンドポイント
	routes.RoutesSetUp(e, db, dbGorm)

	// サーバー開始
	fmt.Println("Server is running on port: ", os.Getenv("PORT"))
	e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
