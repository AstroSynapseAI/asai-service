package gorm

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type GormSeeder struct {
	db *database.Database
	gorm *gorm.DB
	seeders []database.ModelSeeder
}

func NewGormSeeder(db *database.Database) *GormSeeder {
	return &GormSeeder{
		db: db,
		gorm: db.Adapter.Gorm(),
	}
}

func (s *GormSeeder) AddSeeder(seeders ...database.ModelSeeder) *GormSeeder {
	s.seeders = append(s.seeders, seeders...)
	return s
}

func (s *GormSeeder) Run() error {
	for _, seeder := range s.seeders {
		err := seeder.SeedModel(s.db)
		if err != nil {
			fmt.Printf("Failed to seed model: %s\n", err)
			return err
		}
	}
	return nil
}