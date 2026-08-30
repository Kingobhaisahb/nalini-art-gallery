package services

import (
	"os"

	"github.com/razorpay/razorpay-go"
)

type RazorpayService struct {
	Client *razorpay.Client
}

func NewRazorpayService() *RazorpayService {

	keyID := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")

	client := razorpay.NewClient(keyID, keySecret)

	return &RazorpayService{
		Client: client,
	}
}

func (s *RazorpayService) CreateOrder(
	amount float64,
	currency string,
	receipt string,
) (string, error) {

	amountInPaise := int64(amount * 100)

	data := map[string]interface{}{
		"amount":   amountInPaise,
		"currency": currency,
		"receipt":  receipt,
	}

	order, err := s.Client.Order.Create(data, nil)

	if err != nil {
		return "", err
	}

	orderID, ok := order["id"].(string)

	if !ok {
		return "", os.ErrInvalid
	}

	return orderID, nil
}