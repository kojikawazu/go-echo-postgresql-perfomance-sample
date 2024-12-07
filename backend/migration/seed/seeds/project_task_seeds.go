package seeds

import (
	"fmt"
	"log"
	"time"

	project_model "backend/src/models/project"
	task_model "backend/src/models/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedProjectTasks プロジェクトとタスクのシードデータを挿入
func SeedProjectTasks(db *gorm.DB) error {
	log.Println("Seeding project and task tasks...")

	// トランザクションを使用
	return db.Transaction(func(tx *gorm.DB) error {
		assignedTo := "1e678749-1805-4bdc-ad49-315435965458"

		// プロジェクトデータ
		projects := []project_model.Project{
			{
				ID:          uuid.NewString(),
				Name:        "Project Alpha",
				Description: "This is Project Alpha",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          uuid.NewString(),
				Name:        "Project Beta",
				Description: "This is Project Beta",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		// プロジェクトを挿入
		if err := tx.Create(&projects).Error; err != nil {
			return fmt.Errorf("failed to seed projects: %w", err)
		}

		// タスクデータ
		tasks := []task_model.Task{
			{
				ID:          uuid.NewString(),
				ProjectID:   projects[0].ID,
				Name:        "Task 1 for Project Alpha",
				Description: "Description for Task 1",
				Status:      "pending",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				AssignedTo:  &assignedTo,
			},
			{
				ID:          uuid.NewString(),
				ProjectID:   projects[0].ID,
				Name:        "Task 2 for Project Alpha",
				Description: "Description for Task 2",
				Status:      "completed",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				AssignedTo:  &assignedTo,
			},
			{
				ID:          uuid.NewString(),
				ProjectID:   projects[1].ID,
				Name:        "Task 1 for Project Beta",
				Description: "Description for Task 1",
				Status:      "in_progress",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				AssignedTo:  &assignedTo,
			},
		}

		// タスクを挿入
		if err := tx.Create(&tasks).Error; err != nil {
			return fmt.Errorf("failed to seed tasks: %w", err)
		}

		return nil
	})
}
