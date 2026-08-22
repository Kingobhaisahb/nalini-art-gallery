package dto

type CheckoutRequest struct {
	AddressID uint `json:"address_id" binding:"required"`
}

type OrderItemResponse struct {
	ID         uint    `json:"id"`
	PaintingID uint    `json:"painting_id"`
	Price      float64 `json:"price"`
}

type OrderResponse struct {
	ID         uint               `json:"id"`
	UserID     uint               `json:"user_id"`
	AddressID  uint               `json:"address_id"`
	TotalPrice float64            `json:"total_price"`
	Status     string             `json:"status"`
	Items      []OrderItemResponse `json:"items"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=PENDING CONFIRMED SHIPPED DELIVERED CANCELLED"`
}