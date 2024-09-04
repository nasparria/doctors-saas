// internal/domain/patient.go
package domain

import "time"

type Patient struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
