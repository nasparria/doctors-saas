// internal/repository/doctor_repository.go
package repository

import (
	"context"
	"doctors/internal/domain"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	GetDefaultDoctor(ctx context.Context) (*domain.Doctor, error)
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) GetDefaultDoctor(ctx context.Context) (*domain.Doctor, error) {
	var doctor domain.Doctor
	if err := r.db.WithContext(ctx).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}
