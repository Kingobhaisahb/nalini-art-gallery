package models

type OrderItem struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID    uint    `gorm:"not null" json:"order_id"`
	PaintingID uint    `gorm:"not null" json:"painting_id"`
	Price      float64 `gorm:"type:decimal(10,2);not null" json:"price"`

	Painting Painting `gorm:"foreignKey:PaintingID" json:"painting"`
}