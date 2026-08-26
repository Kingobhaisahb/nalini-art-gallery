package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	OrderService services.OrderService
}

func (o *OrderController) Checkout(c *gin.Context) {

	var req dto.CheckoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	userID := userIDValue.(uint)

	order, err := o.OrderService.Checkout(
		userID,
		req.AddressID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Checkout successful",
		"order":   buildOrderResponse(order),
	})
}

func (o *OrderController) GetMyOrders(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	userID := userIDValue.(uint)

	orders, err := o.OrderService.GetUserOrders(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch orders",
		})
		return
	}

	responses := make([]dto.OrderResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(
			responses,
			buildOrderResponse(&order),
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"orders":  responses,
	})
}

func (o *OrderController) GetMyOrder(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid order ID",
		})
		return
	}

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	userID := userIDValue.(uint)

	order, err := o.OrderService.GetUserOrder(
		uint(id),
		userID,
	)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Order not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"order":   buildOrderResponse(order),
	})
}

func buildOrderResponse(
	order *models.Order,
) dto.OrderResponse {

	items := make([]dto.OrderItemResponse, 0)

	for _, item := range order.OrderItems {

		items = append(
			items,
			dto.OrderItemResponse{
				ID:         item.ID,
				PaintingID: item.PaintingID,
				Price:      item.Price,
			},
		)
	}

	return dto.OrderResponse{
		ID:         order.ID,
		UserID:     order.UserID,
		AddressID:  order.AddressID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		Items:      items,
		CreatedAt:  order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (o *OrderController) GetAllOrders(c *gin.Context) {

	orders, err := o.OrderService.GetAllOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch orders",
		})
		return
	}

	responses := make([]dto.OrderResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(
			responses,
			buildOrderResponse(&order),
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"orders":  responses,
	})
}

func (o *OrderController) GetAdminOrder(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid order ID",
		})
		return
	}

	order, err := o.OrderService.GetAdminOrder(uint(id))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Order not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"order":   buildOrderResponse(order),
	})
}

func (o *OrderController) UpdateOrderStatus(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid order ID",
		})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Status is required",
		})
		return
	}

	err = o.OrderService.UpdateOrderStatus(
		uint(id),
		req.Status,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order status updated successfully",
	})
}