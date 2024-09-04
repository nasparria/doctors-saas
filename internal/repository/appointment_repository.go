package repository

import (
	"context"
	"doctors/internal/domain"
	"gorm.io/gorm"
	"time"
)

type AppointmentRepository interface {
	Create(ctx context.Context, appointment *domain.Appointment) error
	GetByID(ctx context.Context, id uint) (*domain.Appointment, error)
	Update(ctx context.Context, appointment *domain.Appointment) error
	Delete(ctx context.Context, id uint) error
	GetByDate(ctx context.Context, date time.Time) ([]domain.Appointment, error)
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(ctx context.Context, appointment *domain.Appointment) error {
	return r.db.WithContext(ctx).Create(appointment).Error
}

func (r *appointmentRepository) GetByID(ctx context.Context, id uint) (*domain.Appointment, error) {
	var appointment domain.Appointment
	err := r.db.WithContext(ctx).First(&appointment, id).Error
	return &appointment, err
}

func (r *appointmentRepository) Update(ctx context.Context, appointment *domain.Appointment) error {
	return r.db.WithContext(ctx).Save(appointment).Error
}

func (r *appointmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Appointment{}, id).Error
}

func (r *appointmentRepository) GetByDate(ctx context.Context, date time.Time) ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	err := r.db.WithContext(ctx).Where("DATE(date_time) = ?", date.Format("2006-01-02")).Find(&appointments).Error
	return appointments, err
}
