package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Email        string    `gorm:"size:255;unique;not null" json:"email"`
	Password     string    `gorm:"size:255" json:"-"`
	Phone        string    `gorm:"size:15;not null" json:"phone"`
	Role         string    `gorm:"type:enum('ADMIN','CUSTOMER');default:CUSTOMER" json:"role"`
	GoogleID     *string   `gorm:"size:255" json:"google_id,omitempty"`
	AuthProvider string    `gorm:"type:enum('LOCAL','GOOGLE');default:LOCAL" json:"auth_provider"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}