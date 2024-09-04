// internal/usecase/appointment_usecase.go
package usecase

import (
	"context"
	"doctors/internal/domain"
	"doctors/internal/repository"
	"doctors/pkg/email"
	"fmt"
	"time"
)

type AppointmentUseCase interface {
	CreateAppointment(ctx context.Context, appointment *domain.Appointment) error
	GetAppointment(ctx context.Context, id uint) (*domain.Appointment, error)
	UpdateAppointment(ctx context.Context, appointment *domain.Appointment) error
	DeleteAppointment(ctx context.Context, id uint) error
	GetAppointmentsByDate(ctx context.Context, date time.Time) ([]domain.Appointment, error)
	SendReminders(ctx context.Context) error
}

type appointmentUseCase struct {
	appointmentRepo repository.AppointmentRepository
	patientRepo     repository.PatientRepository
	doctorRepo      repository.DoctorRepository
	emailSender     email.Sender
}

func NewAppointmentUseCase(
	appointmentRepo repository.AppointmentRepository,
	patientRepo repository.PatientRepository,
	doctorRepo repository.DoctorRepository,
	emailSender email.Sender,
) AppointmentUseCase {
	return &appointmentUseCase{
		appointmentRepo: appointmentRepo,
		patientRepo:     patientRepo,
		doctorRepo:      doctorRepo,
		emailSender:     emailSender,
	}
}

func (uc *appointmentUseCase) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	// Get the default doctor
	defaultDoctor, err := uc.doctorRepo.GetDefaultDoctor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get default doctor: %w", err)
	}
	appointment.DoctorID = defaultDoctor.ID

	// Create the appointment
	if err := uc.appointmentRepo.Create(ctx, appointment); err != nil {
		return fmt.Errorf("failed to create appointment: %w", err)
	}

	// Get the patient
	patient, err := uc.patientRepo.GetByID(ctx, appointment.PatientID)
	if err != nil {
		return fmt.Errorf("failed to get patient: %w", err)
	}

	// Send email
	subject := "Appointment Confirmation"
	body := fmt.Sprintf("Dear %s,\n\nYour appointment with Dr. %s is confirmed for %s.\n\nNotes: %s\n\nBest regards,\nDoctor SaaS Team",
		patient.Name, defaultDoctor.Name, appointment.DateTime.Format(time.RFC1123), appointment.Notes)

	if err := uc.emailSender.Send(patient.Email, subject, body); err != nil {
		// Log the error but don't fail the appointment creation
		fmt.Printf("Failed to send confirmation email: %v\n", err)
	}

	return nil
}

func (uc *appointmentUseCase) GetAppointment(ctx context.Context, id uint) (*domain.Appointment, error) {
	return uc.appointmentRepo.GetByID(ctx, id)
}

func (uc *appointmentUseCase) UpdateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	// Add any business logic here
	return uc.appointmentRepo.Update(ctx, appointment)
}

func (uc *appointmentUseCase) DeleteAppointment(ctx context.Context, id uint) error {
	return uc.appointmentRepo.Delete(ctx, id)
}

func (uc *appointmentUseCase) GetAppointmentsByDate(ctx context.Context, date time.Time) ([]domain.Appointment, error) {
	return uc.appointmentRepo.GetByDate(ctx, date)
}

func (uc *appointmentUseCase) SendReminders(ctx context.Context) error {
	tomorrow := time.Now().AddDate(0, 0, 1)
	appointments, err := uc.GetAppointmentsByDate(ctx, tomorrow)
	if err != nil {
		return err
	}

	for _, apt := range appointments {
		patient, err := uc.patientRepo.GetByID(ctx, apt.PatientID)
		if err != nil {
			continue
		}

		err = uc.emailSender.Send(patient.Email, "Appointment Reminder", "You have an appointment tomorrow at "+apt.DateTime.Format("15:04"))
		if err != nil {
			// Log the error, but continue sending other reminders
			// TODO: Implement proper error logging
		}
	}

	return nil
}
