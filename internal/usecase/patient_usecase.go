// internal/usecase/patient_usecase.go
package usecase

import (
	"context"
	"doctors/internal/domain"
	"doctors/internal/repository"
)

type PatientUseCase interface {
	CreatePatient(ctx context.Context, patient *domain.Patient) error
	GetPatient(ctx context.Context, id uint) (*domain.Patient, error)
	UpdatePatient(ctx context.Context, patient *domain.Patient) error
	DeletePatient(ctx context.Context, id uint) error
	ListPatients(ctx context.Context, page, pageSize int) ([]domain.Patient, int64, error)
}

type patientUseCase struct {
	patientRepo repository.PatientRepository
}

func NewPatientUseCase(patientRepo repository.PatientRepository) PatientUseCase {
	return &patientUseCase{patientRepo: patientRepo}
}

func (uc *patientUseCase) CreatePatient(ctx context.Context, patient *domain.Patient) error {
	// Add any business logic here
	return uc.patientRepo.Create(ctx, patient)
}

func (uc *patientUseCase) GetPatient(ctx context.Context, id uint) (*domain.Patient, error) {
	return uc.patientRepo.GetByID(ctx, id)
}

func (uc *patientUseCase) UpdatePatient(ctx context.Context, patient *domain.Patient) error {
	// Add any business logic here
	return uc.patientRepo.Update(ctx, patient)
}

func (uc *patientUseCase) DeletePatient(ctx context.Context, id uint) error {
	return uc.patientRepo.Delete(ctx, id)
}

func (uc *patientUseCase) ListPatients(ctx context.Context, page, pageSize int) ([]domain.Patient, int64, error) {
	return uc.patientRepo.List(ctx, page, pageSize)
}
