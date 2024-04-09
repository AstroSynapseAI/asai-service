package gorm

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type GormSeeder struct {
	db      *database.Database
	gorm    *gorm.DB
	seeders []database.ModelSeeder
}

func NewGormSeeder(db *database.Database) *GormSeeder {
	return &GormSeeder{
		db:   db,
		gorm: db.Adapter.Gorm(),
	}
}

func (s *GormSeeder) AddSeeder(seeders ...database.ModelSeeder) *GormSeeder {
	s.seeders = append(s.seeders, seeders...)
	return s
}

func (s *GormSeeder) Run() error {
	for _, seeder := range s.seeders {
		for _, action := range seeder.SeedModel(s.db) {
			result := s.gorm.Where("seeder_name = ?", action.ID).First(&DBSeeder{})

			if result.Error == gorm.ErrRecordNotFound {
				err := action.Execute(s.db)

				if err != nil {
					fmt.Printf("Failed to seed model: %s\n", err)
					return err
				}

				if result := s.gorm.Create(&DBSeeder{SeederName: action.ID}); result.Error != nil {
					fmt.Printf("Failed to seed model: %s\n", result.Error)
					return result.Error
				}
			}

		}
	}
	return nil
}
