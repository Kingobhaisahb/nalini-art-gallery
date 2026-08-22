package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService services.CartService
}

func getUserID(c *gin.Context) (uint, bool) {

	value, exists := c.Get("user_id")

	if !exists {
		return 0, false
	}

	userID, ok := value.(uint)

	return userID, ok
}

func (cc *CartController) AddToCart(c *gin.Context) {

	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	var req dto.AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err := cc.CartService.AddToCart(userID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Painting added to cart",
	})
}

func (cc *CartController) GetCart(c *gin.Context) {

	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	cart, total, err := cc.CartService.GetCart(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch cart",
		})
		return
	}

	items := make([]dto.CartItemResponse, 0, len(cart))

	for _, item := range cart {

		items = append(items, dto.CartItemResponse{
			ID:         item.ID,
			PaintingID: item.PaintingID,
			Title:      item.Painting.Title,
			Price:      item.Painting.Price,
			Status:     item.Painting.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"cart": dto.CartResponse{
			Items:      items,
			TotalPrice: total,
		},
	})
}

func (cc *CartController) RemoveFromCart(c *gin.Context) {

	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	paintingID, err := strconv.ParseUint(
		c.Param("painting_id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	err = cc.CartService.RemoveFromCart(
		userID,
		uint(paintingID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to remove painting from cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Painting removed from cart",
	})
}

func (cc *CartController) ClearCart(c *gin.Context) {

	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	err := cc.CartService.ClearCart(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to clear cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cart cleared successfully",
	})
}