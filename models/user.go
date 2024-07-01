package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Username  string    `gorm:"size:255;not null;" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
