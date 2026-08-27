package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService services.UserService
}

// ==========================================
// GET ALL USERS
// GET /api/admin/users
// ==========================================

func (u *UserController) GetAllUsers(c *gin.Context) {

	users, err := u.UserService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch users",
		})
		return
	}

	responses := make([]dto.AdminUserResponse, 0, len(users))

	for _, user := range users {

		responses = append(
			responses,
			dto.AdminUserResponse{
				ID:           user.ID,
				Name:         user.Name,
				Email:        user.Email,
				Phone:        user.Phone,
				Role:         user.Role,
				AuthProvider: user.AuthProvider,
				CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"users":   responses,
	})
}

// ==========================================
// GET USER DETAILS
// GET /api/admin/users/:id
// ==========================================

func (u *UserController) GetUserByID(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid user ID",
		})
		return
	}

	user, orders, err := u.UserService.GetUserByID(uint(id))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch user",
		})
		return
	}

	orderResponses := make(
		[]dto.OrderResponse,
		0,
		len(orders),
	)

	for _, order := range orders {

		items := make(
			[]dto.OrderItemResponse,
			0,
			len(order.OrderItems),
		)

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

		orderResponses = append(
			orderResponses,
			dto.OrderResponse{
				ID:         order.ID,
				UserID:     order.UserID,
				AddressID:  order.AddressID,
				TotalPrice: order.TotalPrice,
				Status:     order.Status,
				Items:      items,
				CreatedAt:  order.CreatedAt.Format(
					"2006-01-02 15:04:05",
				),
				UpdatedAt: order.UpdatedAt.Format(
					"2006-01-02 15:04:05",
				),
			},
		)
	}

	response := dto.AdminUserDetailsResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Role:         user.Role,
		AuthProvider: user.AuthProvider,
		CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Orders:       orderResponses,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    response,
	})
}

// ==========================================
// UPDATE USER ROLE
// PATCH /api/admin/users/:id/role
// ==========================================

func (u *UserController) UpdateUserRole(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid user ID",
		})
		return
	}

	var req dto.UpdateUserRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	adminIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	adminID := adminIDValue.(uint)

	err = u.UserService.UpdateUserRole(
		adminID,
		uint(id),
		req.Role,
	)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User role updated successfully",
	})
}

