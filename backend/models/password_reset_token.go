package models

import "time"

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"size:255;unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}