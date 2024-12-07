package seeds

import (
	"log"

	auth_user_model "backend/src/models/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	log.Println("Seeding users...")

	// ユーザー用のサンプルデータ
	users := []auth_user_model.User{
		{
			ID:       uuid.NewString(),
			Username: "admin",
			Email:    "admin@example.com",
			Password: hashPassword("password123"), // パスワードをハッシュ化
		},
		{
			ID:       uuid.NewString(),
			Username: "user1",
			Email:    "user1@example.com",
			Password: hashPassword("user1pass"),
		},
	}

	// シードデータの挿入
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user: %v", err)
			return err
		}
	}

	log.Println("Seeding users completed!")
	return nil
}

// パスワードのハッシュ化
func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashed)
}
