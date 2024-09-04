// internal/infrastructure/database/postgres.go
package database

import (
	"doctors/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(host, user, password, dbname string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&domain.Patient{}, &domain.Appointment{}, &domain.Doctor{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	// Create a default doctor if not exists
	var doctor domain.Doctor
	result := db.First(&doctor)
	if result.Error == gorm.ErrRecordNotFound {
		defaultDoctor := domain.Doctor{
			Name:  "Nicolas Asparria",
			Email: "mailtrap@demomailtrap.com",
		}
		if err := db.Create(&defaultDoctor).Error; err != nil {
			return nil, fmt.Errorf("failed to create default doctor: %w", err)
		}
	}

	return db, nil
}
