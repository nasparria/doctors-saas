// internal/domain/doctor.go
package domain

import (
	"time"
)

type Doctor struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
