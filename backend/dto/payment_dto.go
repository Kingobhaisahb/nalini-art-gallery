package dto

type CreatePaymentOrderResponse struct {
	OrderID         uint   `json:"order_id"`
	RazorpayOrderID string `json:"razorpay_order_id"`
	Amount          float64 `json:"amount"`
	Currency        string `json:"currency"`
	Status          string `json:"status"`
}

type VerifyPaymentRequest struct {
	OrderID           uint   `json:"order_id" binding:"required"`
	RazorpayPaymentID string `json:"razorpay_payment_id" binding:"required"`
	RazorpaySignature string `json:"razorpay_signature" binding:"required"`
}