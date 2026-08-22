package models

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	AddressID  uint      `gorm:"not null" json:"address_id"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string    `gorm:"type:enum('PENDING','CONFIRMED','SHIPPED','DELIVERED','CANCELLED');default:PENDING" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}