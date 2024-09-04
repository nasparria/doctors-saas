// internal/repository/patient_repository.go
package repository

import (
	"context"
	"doctors/internal/domain"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *domain.Patient) error
	GetByID(ctx context.Context, id uint) (*domain.Patient, error)
	Update(ctx context.Context, patient *domain.Patient) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, page, pageSize int) ([]domain.Patient, int64, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(ctx context.Context, patient *domain.Patient) error {
	return r.db.WithContext(ctx).Create(patient).Error
}

func (r *patientRepository) GetByID(ctx context.Context, id uint) (*domain.Patient, error) {
	var patient domain.Patient
	err := r.db.WithContext(ctx).First(&patient, id).Error
	return &patient, err
}

func (r *patientRepository) Update(ctx context.Context, patient *domain.Patient) error {
	return r.db.WithContext(ctx).Save(patient).Error
}

func (r *patientRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Patient{}, id).Error
}

func (r *patientRepository) List(ctx context.Context, page, pageSize int) ([]domain.Patient, int64, error) {
	var patients []domain.Patient
	var totalCount int64

	offset := (page - 1) * pageSize

	// Count total number of patients
	if err := r.db.WithContext(ctx).Model(&domain.Patient{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve patients with pagination
	err := r.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&patients).Error
	if err != nil {
		return nil, 0, err
	}

	return patients, totalCount, nil
}
