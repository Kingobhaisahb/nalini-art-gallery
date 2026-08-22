package models

import "time"

type Cart struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	PaintingID uint      `gorm:"not null" json:"painting_id"`

	Painting Painting `gorm:"foreignKey:PaintingID" json:"painting"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Cart) TableName() string {
	return "cart"
}