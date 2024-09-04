package domain

import "time"

type Appointment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PatientID uint      `json:"patient_id"`
	DoctorID  uint      `json:"doctor_id"`
	DateTime  time.Time `json:"date_time"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
