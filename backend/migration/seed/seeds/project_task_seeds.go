package seeds

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	project_model "backend/src/models/project"
	task_model "backend/src/models/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedProjectTasks プロジェクトとタスクのシードデータを挿入
func SeedProjectTasks(db *gorm.DB) error {
	log.Println("Seeding project and task tasks...")

	// カスタム乱数生成器を作成
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	assignedTo := os.Getenv("SEED_USER_ID")

	// トランザクションを使用
	return db.Transaction(func(tx *gorm.DB) error {

		// 大量データ用のループ
		for i := 0; i < 100; i++ { // 100プロジェクト
			projectID := uuid.NewString()
			project := project_model.Project{
				ID:          projectID,
				Name:        fmt.Sprintf("Project %d", i+1),
				Description: fmt.Sprintf("This is Project %d", i+1),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			// プロジェクトを挿入
			if err := tx.Create(&project).Error; err != nil {
				return fmt.Errorf("failed to seed projects: %w", err)
			}

			// タスクをプロジェクトごとに生成
			for j := 0; j < 1000; j++ { // 各プロジェクトに1000タスク
				task := task_model.Task{
					ID:          uuid.NewString(),
					ProjectID:   projectID,
					Name:        fmt.Sprintf("Task %d for Project %d", j+1, i+1),
					Description: fmt.Sprintf("Description for Task %d of Project %d", j+1, i+1),
					Status:      []string{"pending", "completed", "in_progress"}[r.Intn(3)], // カスタム乱数生成器を使用
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					AssignedTo:  &assignedTo,
				}

				// タスクを挿入
				if err := tx.Create(&task).Error; err != nil {
					return fmt.Errorf("failed to seed tasks: %w", err)
				}
			}
		}

		return nil
	})
}
