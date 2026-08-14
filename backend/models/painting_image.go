package models

import "time"

type PaintingImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PaintingID uint      `gorm:"not null" json:"painting_id"`
	ImageURL  string    `gorm:"size:500;not null" json:"image_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}