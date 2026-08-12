package models

import "time"

type Painting struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Title       string  `gorm:"size:200;not null" json:"title"`
	Price       float64 `gorm:"not null" json:"price"`
	Description string  `gorm:"type:text" json:"description"`

	Category string `gorm:"size:100;not null" json:"category"`
	Medium   string `gorm:"size:100;not null" json:"medium"`

	Width  float64 `gorm:"not null" json:"width"`
	Height float64 `gorm:"not null" json:"height"`
	Unit   string  `gorm:"size:20;not null;default:'IN'" json:"unit"`

	Featured bool   `gorm:"default:false" json:"featured"`
	Status   string `gorm:"type:enum('AVAILABLE','SOLD');default:AVAILABLE" json:"status"`

	ProcessVideoURL *string `gorm:"size:500" json:"process_video_url,omitempty"`

	Views uint `gorm:"default:0" json:"views"`

	Tags string `gorm:"type:text" json:"tags"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}