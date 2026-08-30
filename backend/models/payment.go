package models

import "time"

type Payment struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	OrderID uint  `gorm:"not null;index" json:"order_id"`
	Order   Order `gorm:"foreignKey:OrderID" json:"-"`

	RazorpayOrderID   string  `gorm:"size:100;unique;not null" json:"razorpay_order_id"`
	RazorpayPaymentID *string `gorm:"size:100" json:"razorpay_payment_id,omitempty"`
	RazorpaySignature  *string `gorm:"size:255" json:"razorpay_signature,omitempty"`

	Amount   float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency string  `gorm:"size:10;not null;default:INR" json:"currency"`

	Status string `gorm:"type:enum('CREATED','SUCCESS','FAILED');default:CREATED" json:"status"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}