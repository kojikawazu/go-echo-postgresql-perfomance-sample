package seeds

import (
	"gorm.io/gorm"
)

// シードデータを挿入
func SeedAll(db *gorm.DB) error {
	// サンプルデータをシード
	// if err := SeedSamples(db); err != nil {
	// 	return err
	// }

	// ユーザーデータをシード
	// if err := SeedUsers(db); err != nil {
	// 	return err
	// }

	// プロジェクトとタスクのシードデータを挿入
	if err := SeedProjectTasks(db); err != nil {
		return err
	}

	return nil
}
