package seeds

import (
	sample_model "backend/src/models/sample"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedSamples(db *gorm.DB) error {
	log.Println("Seeding samples...")

	// サンプルデータの定義
	samples := []sample_model.Sample{
		{
			ID:    uuid.NewString(),
			Name:  "Sample 1",
			Value: 100,
		},
		{
			ID:    uuid.NewString(),
			Name:  "Sample 2",
			Value: 200,
		},
		{
			ID:    uuid.NewString(),
			Name:  "Sample 3",
			Value: 300,
		},
	}

	// シードデータの挿入
	for _, sample := range samples {
		if err := db.Create(&sample).Error; err != nil {
			log.Printf("Failed to seed sample: %v", err)
			return err
		}
	}

	log.Println("Seeding samples completed!")
	return nil
}
