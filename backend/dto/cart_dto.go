package dto

type AddToCartRequest struct {
	PaintingID uint `json:"painting_id" binding:"required"`
}

type CartItemResponse struct {
	ID         uint    `json:"id"`
	PaintingID uint    `json:"painting_id"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
}

type CartResponse struct {
	Items      []CartItemResponse `json:"items"`
	TotalPrice float64            `json:"total_price"`
}