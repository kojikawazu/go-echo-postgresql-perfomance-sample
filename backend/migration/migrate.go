package main

import (
	"fmt"

	project_model "backend/src/models/project"
	sample_model "backend/src/models/sample"
	task_model "backend/src/models/task"
	auth_user_model "backend/src/models/user"

	"gorm.io/gorm"
)

// AutoMigrate マイグレーション用の関数
func SetUpMigration(db *gorm.DB) error {
	fmt.Println("AutoMigrate start...")

	// UUID生成拡張を有効化（PostgreSQL専用）
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		fmt.Println("AutoMigrate error:", err)
		return err
	}

	// テーブルのマイグレーション
	err := db.AutoMigrate(
		&sample_model.Sample{},
		&project_model.Project{},
		&task_model.Task{},
		&auth_user_model.User{},
	)
	if err != nil {
		fmt.Println("AutoMigrate error:", err)
		return err
	}

	fmt.Println("AutoMigrate end...")
	return nil
}
