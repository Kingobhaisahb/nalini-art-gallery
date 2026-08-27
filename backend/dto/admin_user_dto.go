package dto

type AdminUserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	AuthProvider string `json:"auth_provider"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminUserDetailsResponse struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	Phone        string          `json:"phone"`
	Role         string          `json:"role"`
	AuthProvider string          `json:"auth_provider"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	Orders       []OrderResponse `json:"orders"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=ADMIN CUSTOMER"`
}