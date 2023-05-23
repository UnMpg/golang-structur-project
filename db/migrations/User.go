package migrations

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string    `gorm:"type:varchar(255);not null "`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `gorm:"not null"`
	Role             string    `gorm:"type:varchar(255);not null"`
	Provider         string    `gorm:"not null"`
	Photo            string    `gorm:"not null"`
	VerificationCode string
	Verified         bool `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
